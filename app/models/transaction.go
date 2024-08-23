package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	BaseModel
	TransactionID        string  `gorm:"type:uuid" json:"transactionId"`
	FromAccountID        int64   `gorm:"type:bigint not null" json:"fromAccountId"`
	ToAccountID          int64   `gorm:"type:bigint not null" json:"toAccountId"`
	Amount               float64 `gorm:"type:bigint not null" json:"amount"`
	Status               string  `gorm:"type:varchar(100) not null" json:"status"`
	Description          string  `gorm:"type:varchar(255) not null" json:"description"`
	TransactionType      string  `gorm:"type:varchar(100) not null" json:"transactionType"`
	TransactionReference string  `gorm:"type:varchar(255) not null" json:"transactionReference"`
}

type CreateTransferDto struct {
	FromAccountID int64   `json:"fromAccountId"`
	ToAccountID   int64   `json:"toAccountId" validate:"required"`
	Amount        float64 `json:"amount" validate:"required"`
}

// BeforeCreate Hook
func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	generateTransactionID(t)
	return
}

// Generate UUIDs before inserting a new transaction
func generateTransactionID(t *Transaction) error {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	t.TransactionID = uuid.String()
	return nil
}
