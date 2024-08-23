package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bayuscodings/telloservice/app/models"
	"github.com/bayuscodings/telloservice/app/utils"
	"github.com/bayuscodings/telloservice/app/validators"
	"github.com/go-playground/validator/v10"
)

// BaseController is a base structure for all controllers
type BaseController struct {
	Validator *validator.Validate
}

func NewBaseController() *BaseController {
	validator := validator.New()

	// Register the custom validator
	validator.RegisterValidation("currency", validators.ValidCurrency)

	return &BaseController{
		Validator: validator,
	}
}

func (bc *BaseController) asApiResponse(w http.ResponseWriter, message string, data interface{}, code ...int) {
	httpStatusCode := 200
	if len(code) > 0 {
		httpStatusCode = code[0]
	}

	response := models.ApiResponse[interface{}]{
		StatusCode: httpStatusCode,
		Message:    message,
		Data:       data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	json.NewEncoder(w).Encode(response)
}

func (bc *BaseController) asPaginatedApiResponse(
	w http.ResponseWriter,
	message string,
	data interface{},
	pagination models.PaginationOutput,
) {
	httpStatusCode := http.StatusOK

	response := models.PaginatedApiResponse[interface{}]{
		StatusCode: httpStatusCode,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	json.NewEncoder(w).Encode(response)
}

// parsePaginationParams extracts and parses pagination parameters from the request.
func (bc *BaseController) parsePaginationParams(r *http.Request) models.PaginationInput {
	page := utils.ParseInt(r.URL.Query().Get("page"), 1)
	limit := utils.ParseInt(r.URL.Query().Get("limit"), 20)
	orderBy := models.ParseNullString(r.URL.Query().Get("orderBy"))
	orderDir := models.ParseNullString(r.URL.Query().Get("orderDir"))

	return models.PaginationInput{
		Page:     page,
		Limit:    limit,
		OrderBy:  orderBy,
		OrderDir: orderDir,
	}
}
