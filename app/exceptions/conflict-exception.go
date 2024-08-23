package exceptions

import "net/http"

type ConflictException struct {
	*BaseHTTPException
}

func NewConflictException(message string) *ConflictException {
	return &ConflictException{
		BaseHTTPException: NewHTTPException(http.StatusConflict, message),
	}
}
