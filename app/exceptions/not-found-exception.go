package exceptions

import "net/http"

type NotFoundException struct {
	*BaseHTTPException
}

func NewNotFoundException(message string) *NotFoundException {
	return &NotFoundException{
		BaseHTTPException: NewHTTPException(http.StatusNotFound, message),
	}
}
