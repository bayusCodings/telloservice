package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/bayuscodings/telloservice/app/auth"
	"github.com/bayuscodings/telloservice/app/exceptions"
	"github.com/bayuscodings/telloservice/app/models"
	"github.com/bayuscodings/telloservice/app/repositories"
	"github.com/bayuscodings/telloservice/app/utils"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type UserService struct {
	repo *repositories.UserRepository
	JWT  auth.TokenMaker
}

// NewUserService creates a new instance of UserService
func NewUserService(DB *gorm.DB, JWT auth.TokenMaker) *UserService {
	repo := repositories.NewUserRepository(DB)
	return &UserService{
		repo: repo,
		JWT:  JWT,
	}
}

// CreateUser creates a new user in the database
func (s *UserService) CreateUser(input *models.CreateUserInputDto) (*models.UserResponseDto, exceptions.HTTPException) {
	existingUser, err := s.repo.FindByUsernameOrEmail(input.Username, input.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Msg(fmt.Sprintf("Failed to check existing user: %v", err))
		return nil, exceptions.NewInternalServerException("Error occured while creating user")
	}

	if existingUser != nil {
		log.Printf("user already exists")
		return nil, exceptions.NewConflictException("user already exists")
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return nil, exceptions.NewInternalServerException("Error occured while creating user")
	}

	user := models.User{
		Username:  input.Username,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     strings.ToLower(input.Email),
		Password:  hashedPassword,
	}

	if err := s.repo.Create(&user); err != nil {
		log.Error().Msg(fmt.Sprintf("Failed to create user: %v", err))
		return nil, exceptions.NewInternalServerException("Failed to create user")
	}

	userResponse := models.ToUserResponseDto(user)
	return &userResponse, nil
}

// Authenticate User
func (s *UserService) Login(input *models.UserLoginInputDto) (*models.LoginResponseDto, exceptions.HTTPException) {
	user, err := s.repo.FindByEmail(input.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewNotFoundException(fmt.Sprintf("User with email: %s does not exist", input.Email))
		}

		log.Error().Msg(fmt.Sprintf("Error retrieving user by email %v", err))
		return nil, exceptions.NewInternalServerException("Error occured while authenticating user")
	}

	err = utils.CheckPassword(input.Password, user.Password)
	if err != nil {
		return nil, exceptions.NewUnauthorizedException("Invalid password")
	}

	duration := time.Hour * 2
	token, err := s.JWT.CreateToken(user.ID, user.Email, duration)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Error generating jwt token %v", err))
		return nil, exceptions.NewInternalServerException("Error occured while authenticating user")
	}

	result := &models.LoginResponseDto{
		Token: token,
	}

	return result, nil
}

// FetchUserById retrieves a user by id from the database
func (s *UserService) FetchUserById(id uint) (*models.UserResponseDto, exceptions.HTTPException) {
	user, err := s.repo.FindById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("User with id: %d does not exist", id)
			return nil, exceptions.NewNotFoundException("User not found")
		}

		log.Error().Msg(fmt.Sprintf("Error retrieving user by id %v", err))
		return nil, exceptions.NewInternalServerException("Error occured while retrieving user")
	}

	userResponse := models.ToUserResponseDto(*user)
	return &userResponse, nil
}
