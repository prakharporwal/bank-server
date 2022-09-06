package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token is expired")
)

type JWTMaker struct {
	SecretKey string
}

const minSecretKeySize = 32

func NewJWTMaker(secretKey string) (TokenMaker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

func (jwtMaker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload := NewPayload(username, duration)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(jwtMaker.SecretKey))
}

func (jwtMaker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(jwtToken *jwt.Token) (interface{}, error) {
		_, ok := jwtToken.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(jwtMaker.SecretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		ver, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(ver.Inner, ErrExpiredToken) {
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
