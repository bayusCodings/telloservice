package exceptions

import "net/http"

type UnprocessableEntityException struct {
	*BaseHTTPException
}

func NewUnprocessableEntityException(message string) *UnprocessableEntityException {
	return &UnprocessableEntityException{
		BaseHTTPException: NewHTTPException(http.StatusUnprocessableEntity, message),
	}
}
