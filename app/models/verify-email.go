package models

import "time"

type VerifyEmail struct {
	BaseModel
	Username   string    `gorm:"type:varchar(255) not null" json:"username"`
	Email      string    `gorm:"type:varchar(255) not null" json:"email"`
	SecretCode string    `gorm:"type:varchar(255) not null" json:"secretCode"`
	IsUsed     bool      `gorm:"type:boolean not null default:false" json:"isUsed"`
	ExpiredAt  time.Time `json:"expiredAt"`
}
