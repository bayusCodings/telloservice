package models

type Account struct {
	BaseModel
	UserID   int64   `gorm:"type:bigint not null" json:"userId"`
	Balance  float64 `gorm:"type:bigint not null" json:"balance"`
	Currency string  `gorm:"type:varchar(50) not null" json:"currency"`
}

type CreateAccountInputDto struct {
	UserID   int64  `json:"userId"`
	Currency string `json:"currency" validate:"required,currency"`
}
