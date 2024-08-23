package exceptions

import "net/http"

type BadRequestException struct {
	*BaseHTTPException
}

func NewBadRequestException(message string) *BadRequestException {
	return &BadRequestException{
		BaseHTTPException: NewHTTPException(http.StatusBadRequest, message),
	}
}
