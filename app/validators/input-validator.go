package validators

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/bayuscodings/telloservice/app/exceptions"
	"github.com/bayuscodings/telloservice/app/middlewares"
	"github.com/go-playground/validator/v10"
)

func ValidateInput(validator *validator.Validate, dto interface{}) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Create a new instance of the expected struct type
			request := reflect.New(reflect.TypeOf(dto)).Interface()

			// Decode the request body into the struct
			if jsonError := json.NewDecoder(r.Body).Decode(request); jsonError != nil {
				exception := exceptions.NewBadRequestException("Invalid request payload")
				exception.Respond(w)
				return
			}

			if !ValidateAndRespond(validator, w, request) {
				return
			}

			// Store the validated request in the context for the next handler to use
			ctx := r.Context()
			ctx = context.WithValue(ctx, middlewares.ValidatedRequestKey, request)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
