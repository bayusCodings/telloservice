package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/bayuscodings/telloservice/app/auth"
	"github.com/bayuscodings/telloservice/app/exceptions"
)

func AuthMiddleware(JWT auth.TokenMaker) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the token from the Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				exception := exceptions.NewUnauthorizedException("Authorization header is missing")
				exception.Respond(w)
				return
			}

			// Bearer token format: "Bearer <token>"
			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
				exception := exceptions.NewUnauthorizedException("Invalid Authorization header format")
				exception.Respond(w)
				return
			}

			token := tokenParts[1]

			// Verify the token
			payload, err := JWT.VerifyToken(token)
			if err != nil {
				if err == auth.ErrExpiredToken {
					exception := exceptions.NewUnauthorizedException("Token has expired")
					exception.Respond(w)
				} else {
					exception := exceptions.NewUnauthorizedException("Invalid token")
					exception.Respond(w)
				}
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, UserPayloadKey, payload)

			// Proceed to the next handler
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
