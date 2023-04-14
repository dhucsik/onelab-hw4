package service

import (
	"context"
	"errors"

	"github.com/dhucsik/onelab-hw4/models"
	"github.com/dhucsik/onelab-hw4/storage"
)

type ITransactionService interface {
	Create(ctx context.Context, transaction *models.Transaction) (string, error)
	Update(ctx context.Context, ID string, transaction *models.Transaction) error
	Delete(ctx context.Context, ID string) error
	Get(ctx context.Context, ID string) (models.Transaction, error)
	List(ctx context.Context) ([]models.Transaction, error)
}

type Manager struct {
	Transaction ITransactionService
}

func NewManager(storage *storage.Storage) (*Manager, error) {
	if storage == nil {
		return nil, errors.New("no storage provided")
	}

	tSrv := NewTransactionService(storage)

	return &Manager{
		Transaction: tSrv,
	}, nil
}
