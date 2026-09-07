package account

import (
	"context"
	"errors"
	"time"
)

var ErrNotFound = errors.New("account not found")
var ErrInsufficientFunds = errors.New("insufficient funds")

type Service interface {
	Create(context.Context, *Account) (*Account, error)
	Get(context.Context, string) (*Account, error)
	Debit(context.Context, string, float64) (float64, error)
}
type service struct{ repo Repository }

func NewService(repo Repository) Service { return &service{repo: repo} }

func (s *service) Create(ctx context.Context, account *Account) (*Account, error) {
	if err := s.repo.Create(ctx, account); err != nil {
		return nil, err
	}
	return account, nil
}

func (s *service) Get(ctx context.Context, id string) (*Account, error) {
	a, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if a == nil {
		return nil, ErrNotFound
	}
	return a, nil
}
func (s *service) Debit(ctx context.Context, id string, amount float64) (float64, error) {
	a, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return 0, err
	}
	if a == nil {
		return 0, ErrNotFound
	}

	time.Sleep(50 * time.Millisecond)
	if a.Balance < amount {
		return 0, ErrInsufficientFunds
	}
	newBalance := a.Balance - amount
	if err := s.repo.SetBalance(ctx, id, newBalance); err != nil {
		return 0, err
	}
	return newBalance, nil
}
