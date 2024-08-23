package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenMaker interface {
	CreateToken(id int64, email string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) TokenMaker {
	return &JWTMaker{secretKey}
}

func (maker *JWTMaker) CreateToken(id int64, email string, duration time.Duration) (string, error) {
	payload := NewPayload(id, email, duration)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, err
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		result, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(result.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
