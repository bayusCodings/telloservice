package exceptions

import (
	"encoding/json"
	"net/http"
)

type HTTPException interface {
	GetStatusCode() int
	GetMessage() string
	Respond(w http.ResponseWriter)
}

type BaseHTTPException struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func (e *BaseHTTPException) GetStatusCode() int {
	return e.StatusCode
}

func (e *BaseHTTPException) GetMessage() string {
	return e.Message
}

func (e *BaseHTTPException) Respond(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.GetStatusCode())
	json.NewEncoder(w).Encode(e)
}

func NewHTTPException(code int, message string) *BaseHTTPException {
	return &BaseHTTPException{
		StatusCode: code,
		Message:    message,
	}
}
