package storage

import (
	"context"

	"github.com/dhucsik/onelab-hw4/config"
	"github.com/dhucsik/onelab-hw4/models"
	"github.com/dhucsik/onelab-hw4/storage/postgre"
)

type ITransactionRepository interface {
	Create(ctx context.Context, transaction *models.Transaction) (string, error)
	Get(ctx context.Context, ID string) (models.Transaction, error)
	Update(ctx context.Context, ID string, transaction *models.Transaction) error
	List(ctx context.Context) ([]models.Transaction, error)
	Delete(ctx context.Context, ID string) error
}

type Storage struct {
	Transaction ITransactionRepository
}

func New(ctx context.Context, cfg *config.Config) (*Storage, error) {
	pgDB, err := postgre.Dial(ctx, cfg.PgURL)
	if err != nil {
		return nil, err
	}

	transactionRepo := postgre.NewTransactionRepository(pgDB)

	storage := Storage{
		Transaction: transactionRepo,
	}
	return &storage, nil
}
