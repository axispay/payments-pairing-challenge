package payment

import (
	"context"
	"encoding/json"
	"time"
)

type EventPublisher interface {
	Publish(context.Context, string, []byte) error
}
type Service interface {
	Create(context.Context, string, float64, string) (*Payment, error)
}
type service struct {
	repo      Repository
	publisher EventPublisher
}

func NewService(repo Repository, p EventPublisher) Service { return &service{repo: repo, publisher: p} }
func (s *service) Create(ctx context.Context, accountID string, amount float64, currency string) (*Payment, error) {
	p := &Payment{AccountID: accountID, Amount: amount, Currency: currency, Status: "charged", CreatedAt: time.Now()}
	if err := s.repo.Create(ctx, p); err != nil {
		return nil, err
	}
	payload, _ := json.Marshal(map[string]any{"type": "PAYMENT_CREATED", "payment": p})
	if err := s.publisher.Publish(ctx, "payment-events", payload); err != nil {
		return nil, err
	}
	return p, nil
}
