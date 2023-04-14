package service

import (
	"context"
	"time"

	"github.com/dhucsik/onelab-hw4/models"
	"github.com/dhucsik/onelab-hw4/storage"
)

type TransactionService struct {
	repo *storage.Storage
}

func NewTransactionService(repo *storage.Storage) *TransactionService {
	return &TransactionService{
		repo: repo,
	}
}

func (s *TransactionService) Create(ctx context.Context, transaction *models.Transaction) (string, error) {
	transaction.PaymentTime = time.Now()

	return s.repo.Transaction.Create(ctx, transaction)
}

func (s *TransactionService) Update(ctx context.Context, ID string, transaction *models.Transaction) error {
	return s.repo.Transaction.Update(ctx, ID, transaction)
}

func (s *TransactionService) Delete(ctx context.Context, ID string) error {
	return s.repo.Transaction.Delete(ctx, ID)
}

func (s *TransactionService) Get(ctx context.Context, ID string) (models.Transaction, error) {
	return s.repo.Transaction.Get(ctx, ID)
}

func (s *TransactionService) List(ctx context.Context) ([]models.Transaction, error) {
	return s.repo.Transaction.List(ctx)
}
