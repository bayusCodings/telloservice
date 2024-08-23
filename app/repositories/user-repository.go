package repositories

import (
	"github.com/bayuscodings/telloservice/app/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	*Repository[models.User]
}

func NewUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{
		Repository: New(DB, models.User{}),
	}
}

// Create a new user
func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

// Find a user by ID
func (r *UserRepository) FindById(id uint) (*models.User, error) {
	user := r.Entity
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Find a user by email
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	user := r.Entity
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Find a user by username or email
func (r *UserRepository) FindByUsernameOrEmail(username string, email string) (*models.User, error) {
	user := r.Entity
	if err := r.DB.Where("username = ? OR email = ?", username, email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Update an existing user
func (r *UserRepository) Update(user *models.User) error {
	return r.DB.Save(user).Error
}

// Delete a user by ID
func (r *UserRepository) Delete(id uint) error {
	return r.DB.Delete(&r.Entity, id).Error
}
