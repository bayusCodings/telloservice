package models

import (
	"time"
)

type User struct {
	BaseModel
	Username          string    `gorm:"type:varchar(255) not null unique" json:"username"`
	Password          string    `gorm:"type:varchar(255) not null" json:"password"`
	FirstName         string    `gorm:"type:varchar(150) not null" json:"firstName"`
	LastName          string    `gorm:"type:varchar(150) not null" json:"lastName"`
	Email             string    `gorm:"type:varchar(255) not null unique" json:"email"`
	IsEmailVerified   bool      `gorm:"default:false" json:"isEmailVerified"`
	PasswordChangedAt time.Time `json:"passwordChangedAt"`
	Accounts          []Account `gorm:"foreignKey:UserID" json:"accounts"`
}

type CreateUserInputDto struct {
	Username  string `json:"username" validate:"required,min=3,max=30"`
	Password  string `json:"password" validate:"required,min=6"`
	FirstName string `json:"firstName" validate:"required,alpha,min=2,max=50"`
	LastName  string `json:"lastName" validate:"required,alpha,min=2,max=50"`
	Email     string `json:"email" validate:"required,email"`
}

type UserLoginInputDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserResponseDto struct {
	ID              int64     `json:"id"`
	Username        string    `json:"username"`
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	Email           string    `json:"email"`
	IsEmailVerified bool      `json:"isEmailVerified"`
	CreatedAt       time.Time `json:"createdAt"`
}

type LoginResponseDto struct {
	Token string `json:"token"`
}

func ToUserResponseDto(user User) UserResponseDto {
	return UserResponseDto{
		ID:              user.ID,
		Username:        user.Username,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Email:           user.Email,
		IsEmailVerified: user.IsEmailVerified,
		CreatedAt:       user.CreatedAt,
	}
}
