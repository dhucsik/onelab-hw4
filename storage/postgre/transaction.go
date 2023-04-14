package postgre

import (
	"context"
	"strconv"
	"time"

	"github.com/dhucsik/onelab-hw4/models"
	"gorm.io/gorm"
)

type Transaction struct {
	ID            uint           `gorm:"primaryKey"`
	CreatedAt     time.Time      ``
	UpdatedAt     time.Time      ``
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	SenderID      uint
	PaymentTime   time.Time
	ItemID        uint
	PaymentAmount float64
}

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(ctx context.Context, transaction *models.Transaction) (string, error) {
	model, err := toPostgreTransaction(transaction)
	if err != nil {
		return "", err
	}

	result := r.db.WithContext(ctx).Omit("deleted_at").Create(&model)
	return strconv.FormatUint(uint64(model.ID), 10), result.Error
}

func (r *TransactionRepository) Update(ctx context.Context, ID string, transaction *models.Transaction) error {
	id, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		return err
	}

	model, err := toPostgreTransaction(transaction)
	if err != nil {
		return err
	}

	model.ID = uint(id)
	return r.db.Save(&model).Error
}

func (r *TransactionRepository) Get(ctx context.Context, ID string) (models.Transaction, error) {
	transaction := new(Transaction)
	err := r.db.WithContext(ctx).Where("id = ?", ID).First(transaction).Error
	if err != nil {
		return models.Transaction{}, err
	}

	return toTransactionModel(transaction), nil
}

func (r *TransactionRepository) Delete(ctx context.Context, ID string) error {
	return r.db.WithContext(ctx).Delete(&Transaction{}, ID).Error
}

func (r *TransactionRepository) List(ctx context.Context) ([]models.Transaction, error) {
	var transactions []Transaction
	err := r.db.WithContext(ctx).Find(&transactions)
	return toTransactionModels(transactions), err.Error
}

func toPostgreTransaction(t *models.Transaction) (Transaction, error) {
	senderID, err := strconv.ParseUint(t.SenderID, 10, 32)
	if err != nil {
		return Transaction{}, err
	}

	itemID, err := strconv.ParseUint(t.ItemID, 10, 32)
	if err != nil {
		return Transaction{}, err
	}

	return Transaction{
		SenderID:      uint(senderID),
		PaymentTime:   t.PaymentTime,
		ItemID:        uint(itemID),
		PaymentAmount: t.PaymentAmount,
	}, nil
}

func toTransactionModel(t *Transaction) models.Transaction {
	return models.Transaction{
		ID:            strconv.FormatUint(uint64(t.ID), 10),
		SenderID:      strconv.FormatUint(uint64(t.SenderID), 10),
		PaymentTime:   t.PaymentTime,
		ItemID:        strconv.FormatUint(uint64(t.ItemID), 10),
		PaymentAmount: t.PaymentAmount,
	}
}

func toTransactionModels(transactions []Transaction) []models.Transaction {
	out := make([]models.Transaction, len(transactions))

	for i, transaction := range transactions {
		out[i] = toTransactionModel(&transaction)
	}

	return out
}
