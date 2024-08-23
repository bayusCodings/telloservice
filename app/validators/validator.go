package validators

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateAndRespond(v *validator.Validate, w http.ResponseWriter, data interface{}) bool {
	if err := v.Struct(data); err != nil {
		validationErrors := TranslateValidationErrors(err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Validation failed",
			"error":   validationErrors,
		})
		return false
	}
	return true
}

func TranslateValidationErrors(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			tag := fieldError.Tag()
			field := toJSONFieldName(fieldError.Field())

			var errorMessage string
			switch tag {
			case "required":
				errorMessage = fmt.Sprintf("%s is required", field)
			case "min":
				errorMessage = fmt.Sprintf("%s must be at least %s characters long", field, fieldError.Param())
			case "max":
				errorMessage = fmt.Sprintf("%s must be at most %s characters long", field, fieldError.Param())
			case "alpha":
				errorMessage = fmt.Sprintf("%s can only contain alphabetic characters", field)
			case "email":
				errorMessage = fmt.Sprintf("%s must be a valid email address", field)
			case "currency":
				errorMessage = fmt.Sprintf("%s is not valid; must be one of the following values: USD, EUR, NGN", field)
			default:
				errorMessage = fmt.Sprintf("%s is not valid", field)
			}

			errors[field] = errorMessage
		}
	}

	return errors
}

func toJSONFieldName(fieldName string) string {
	// Convert first letter to lowercase to match JSON field naming convention
	return strings.ToLower(fieldName[:1]) + fieldName[1:]
}
