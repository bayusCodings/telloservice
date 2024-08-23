package repositories

import (
	"gorm.io/gorm"
)

type Repository[T any] struct {
	DB     *gorm.DB
	Entity T
}

func New[T any](DB *gorm.DB, Entity T) *Repository[T] {
	e := &Repository[T]{DB, Entity}
	return e
}
