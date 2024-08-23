package services

import (
	"fmt"

	"github.com/bayuscodings/telloservice/app/exceptions"
	"github.com/bayuscodings/telloservice/app/models"
	"github.com/bayuscodings/telloservice/app/repositories"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type AccountService struct {
	repo *repositories.AccountRepository
}

// NewUserService creates a new instance of AccountService
func NewAccountService(DB *gorm.DB) *AccountService {
	repo := repositories.NewAccountRepository(DB)
	return &AccountService{
		repo: repo,
	}
}

// CreateUser creates a new user in the database
func (s *AccountService) CreateAccount(input *models.CreateAccountInputDto) (*models.Account, exceptions.HTTPException) {
	account := models.Account{
		UserID:   input.UserID,
		Currency: input.Currency,
		Balance:  0,
	}

	if err := s.repo.Create(&account); err != nil {
		// Check for unique violation error
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return nil, exceptions.NewConflictException("User account with this currency already exists")
		}

		log.Error().Msg(fmt.Sprintf("Failed to create account: %v", err))
		return nil, exceptions.NewInternalServerException("Failed to create account")
	}

	return &account, nil
}

// FetchAccounts fetches accounts with pagination
func (s *AccountService) FetchAccounts(
	userId uint,
	paginationInput *models.PaginationInput,
) (*models.PaginatedData[models.Account], exceptions.HTTPException) {
	models.PreparePagination(paginationInput)

	accounts, total, err := s.repo.FetchAccountsWithPagination(userId, paginationInput)
	if err != nil {
		return nil, exceptions.NewInternalServerException("Failed to fetch accounts")
	}

	paginatedData := &models.PaginatedData[models.Account]{}
	return paginatedData.PreparePaginatedResponse(
		&accounts,
		total,
		paginationInput.Limit,
		paginationInput.Page,
	), nil
}

func (s *AccountService) TransferFunds(userId int64, input *models.CreateTransferDto) exceptions.HTTPException {
	sender, err := s.repo.FindUserAccount(uint(input.FromAccountID), uint(userId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return exceptions.NewNotFoundException(fmt.Sprintf("User does not have accountId: %d", input.FromAccountID))
		}

		log.Printf("An error occured while processing transfer: %v", err)
		return exceptions.NewInternalServerException("An error occured while processing transfer")
	}

	if sender.Balance < input.Amount {
		return exceptions.NewBadRequestException("Insufficient balance")
	}

	receiver, err := s.repo.FindById(uint(input.ToAccountID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return exceptions.NewNotFoundException(fmt.Sprintf("toAccountId: %d does not exist", input.ToAccountID))
		}

		log.Error().Msg(fmt.Sprintf("An error occured while processing transfer: %v", err))
		return exceptions.NewInternalServerException("An error occured while processing transfer")
	}

	if err := s.repo.TransferFunds(sender, receiver, input.Amount); err != nil {
		return exceptions.NewInternalServerException("An error occured while processing transfer")
	}

	return nil
}
