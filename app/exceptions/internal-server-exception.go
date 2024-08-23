package exceptions

import "net/http"

type InternalServerException struct {
	*BaseHTTPException
}

func NewInternalServerException(message string) *InternalServerException {
	return &InternalServerException{
		BaseHTTPException: NewHTTPException(http.StatusInternalServerError, message),
	}
}
