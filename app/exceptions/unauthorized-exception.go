package exceptions

import "net/http"

type UnauthorizedException struct {
	*BaseHTTPException
}

func NewUnauthorizedException(message string) *UnauthorizedException {
	return &UnauthorizedException{
		BaseHTTPException: NewHTTPException(http.StatusUnauthorized, message),
	}
}
