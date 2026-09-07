package analytics

import (
	"context"
	"time"
)

type Service interface {
	CreateTransaction(context.Context, *Transaction) (*Transaction, error)
	DailyTotals(context.Context, string, time.Time, time.Time) (DailyTotals, error)
}
type service struct{ repo Repository }

func NewService(r Repository) Service { return &service{r} }

func (s *service) CreateTransaction(ctx context.Context, transaction *Transaction) (*Transaction, error) {
	if transaction.CreatedAt.IsZero() {
		transaction.CreatedAt = time.Now().UTC()
	}
	if err := s.repo.CreateTransaction(ctx, transaction); err != nil {
		return nil, err
	}
	return transaction, nil
}

func (s *service) DailyTotals(ctx context.Context, id string, start, end time.Time) (DailyTotals, error) {
	txs, err := s.repo.FindTransactions(ctx, id, start, end)
	if err != nil {
		return DailyTotals{}, err
	}
	result := DailyTotals{ByStatus: map[string]int{}}
	for _, tx := range txs {
		if a, ok := tx["amount"].(float64); ok {
			result.Total += a
		} else if a, ok := tx["amount"].(int32); ok {
			result.Total += float64(a)
		}
		result.Count++
		if st, ok := tx["status"].(string); ok && st != "" {
			result.ByStatus[st]++
		}
	}
	return result, nil
}
