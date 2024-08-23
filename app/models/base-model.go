package models

import (
	"database/sql"
	"time"
)

// NullString is a wrapper around sql.NullString
type NullString struct {
	sql.NullString
}
type NullInt64 struct {
	sql.NullInt64
}

// NullString
func ParseNullString(valueStr string) NullString {
	return NullString{
		NullString: sql.NullString{String: valueStr, Valid: valueStr != ""},
	}
}

type BaseModel struct {
	ID        int64     `gorm:"primarykey type:bigserial" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type DateRangeFilter struct {
	StartDate *time.Time `json:"startDate"`
	EndDate   *time.Time `json:"endDate"`
}

type PaginationInput struct {
	Page     int        `json:"page" default:"1"`
	Limit    int        `json:"limit" default:"50"`
	OrderBy  NullString `json:"orderBy"`
	OrderDir NullString `json:"orderDir"`
}

func PreparePagination(paginationInput *PaginationInput) {
	if paginationInput.Limit < 1 {
		paginationInput.Limit = 20
	}
	if paginationInput.Page < 1 {
		paginationInput.Page = 1
	}
	if !paginationInput.OrderBy.Valid {
		paginationInput.OrderBy.Scan("created_at")
	}
	if !paginationInput.OrderDir.Valid {
		paginationInput.OrderDir.Scan("DESC")
	}
}

type PaginationOutput struct {
	Page  int   `json:"page"`
	Pages int64 `json:"pages"`
	Limit int   `json:"limit"`
	Size  int   `json:"size"`
	Total int64 `json:"total"`
}

type PaginatedData[T any] struct {
	Data       []T              `json:"data"`
	Pagination PaginationOutput `json:"pagination"`
}

func (paginatedData *PaginatedData[T]) PreparePaginatedResponse(
	data *[]T, total int64, limit int, page int,
) *PaginatedData[T] {
	var pages int64 = int64(total / int64(limit))
	if pages < 1 {
		pages = 1
	}

	paginatedData.Data = *data
	paginatedData.Pagination = PaginationOutput{
		Page:  page,
		Limit: limit,
		Total: total,
		Size:  len(*data),
		Pages: pages,
	}

	return paginatedData
}

type ApiResponse[T any] struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       T      `json:"data"`
}

type PaginatedApiResponse[T any] struct {
	StatusCode int              `json:"statusCode"`
	Message    string           `json:"message"`
	Data       T                `json:"data"`
	Pagination PaginationOutput `json:"pagination"`
}
