package transfer

import (
	"context"
	"encoding/json"
	"errors"
	"time"
)

var ErrInsufficientFunds = errors.New("insufficient funds")

type EventPublisher interface {
	Publish(context.Context, string, []byte) error
}
type Service interface {
	Transfer(context.Context, string, string, float64) (float64, error)
}
type service struct {
	repo      Repository
	publisher EventPublisher
}

func NewService(repo Repository, p EventPublisher) Service { return &service{repo: repo, publisher: p} }
func (s *service) Transfer(ctx context.Context, fromID, toID string, amount float64) (float64, error) {
	from, err := s.repo.FindAccount(ctx, fromID)
	if err != nil {
		return 0, err
	}
	to, err := s.repo.FindAccount(ctx, toID)
	if err != nil {
		return 0, err
	}
	if from.Balance < amount {
		return 0, ErrInsufficientFunds
	}
	if err := s.repo.SetBalance(ctx, fromID, from.Balance-amount); err != nil {
		return 0, err
	}
	if err := s.repo.SetBalance(ctx, toID, to.Balance+amount); err != nil {
		return 0, err
	}
	if err := s.repo.CreateTransfer(ctx, &Transfer{FromAccountID: fromID, ToAccountID: toID, Amount: amount, CreatedAt: time.Now()}); err != nil {
		return 0, err
	}
	payload, _ := json.Marshal(map[string]any{"fromAccountId": fromID, "toAccountId": toID, "amount": amount})
	_ = s.publisher.Publish(ctx, "transfer-events", payload)
	return from.Balance - amount, nil
}
