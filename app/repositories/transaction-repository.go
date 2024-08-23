package repositories

import (
	"github.com/bayuscodings/telloservice/app/models"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	*Repository[models.Transaction]
}

func NewTransactionRepository(DB *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		Repository: New(DB, models.Transaction{}),
	}
}

func (r *TransactionRepository) Create(transaction *models.Transaction) error {
	return r.DB.Create(transaction).Error
}
