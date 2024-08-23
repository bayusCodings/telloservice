package repositories

import (
	"fmt"

	"github.com/bayuscodings/telloservice/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountRepository struct {
	*Repository[models.Account]
}

func NewAccountRepository(DB *gorm.DB) *AccountRepository {
	return &AccountRepository{
		Repository: New(DB, models.Account{}),
	}
}

// Create a new account
func (r *AccountRepository) Create(account *models.Account) error {
	return r.DB.Create(account).Error
}

func (r *AccountRepository) FindById(id uint) (*models.Account, error) {
	account := r.Entity
	if err := r.DB.First(&account, id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepository) FindUserAccount(id, userId uint) (*models.Account, error) {
	account := r.Entity
	if err := r.DB.Where("id = ? AND user_id = ?", id, userId).First(&account, id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

// FetchAccountsWithPagination retrieves accounts with pagination
func (r *AccountRepository) FetchAccountsWithPagination(
	userId uint,
	paginationInput *models.PaginationInput,
) ([]models.Account, int64, error) {
	var accounts []models.Account
	var total int64

	page := paginationInput.Page
	limit := paginationInput.Limit
	orderBy := paginationInput.OrderBy.String
	orderDir := paginationInput.OrderDir.String

	// Fetch total count
	if err := r.DB.Model(&r.Entity).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination and sorting
	query := r.DB.Model(&r.Entity).Where("user_id = ?", userId)
	if orderBy != "" {
		query = query.Order(orderBy + " " + orderDir)
	}

	err := query.Offset((page - 1) * limit).Limit(limit).Find(&accounts).Error
	return accounts, total, err
}

// Transfer funds between two accounts within a transaction
func (r *AccountRepository) TransferFunds(sender, receiver *models.Account, amount float64) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		reference, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		// Debit the sender's account
		sender.Balance -= amount
		if err := tx.Save(&sender).Error; err != nil {
			return err
		}

		debitTransaction := models.Transaction{
			FromAccountID:        sender.ID,
			ToAccountID:          receiver.ID,
			Amount:               amount,
			Status:               "success",
			TransactionType:      "debit",
			TransactionReference: reference.String(),
			Description:          "Transfer to Account " + fmt.Sprint(receiver.ID),
		}
		if err := tx.Create(&debitTransaction).Error; err != nil {
			return err
		}

		// Credit the receiver's account
		receiver.Balance += amount
		if err := tx.Save(&receiver).Error; err != nil {
			return err
		}

		creditTransaction := models.Transaction{
			FromAccountID:        sender.ID,
			ToAccountID:          receiver.ID,
			Amount:               amount,
			Status:               "success",
			TransactionType:      "credit",
			TransactionReference: reference.String(),
			Description:          "Received from Account " + fmt.Sprint(sender.ID),
		}
		if err := tx.Create(&creditTransaction).Error; err != nil {
			return err
		}

		// Commit the transaction
		return nil
	})
}
