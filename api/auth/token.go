package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

type Payload struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (payload *Payload) Valid() error {
	if payload.ExpiresAt.Before(time.Now()) {
		return jwt.ErrTokenExpired
	}
	return nil
}

func NewPayload(username string, duration time.Duration) *Payload {
	id, _ := uuid.NewRandom()
	currTime := time.Now()

	return &Payload{
		Id:        id,
		Username:  username,
		IssuedAt:  currTime,
		ExpiresAt: currTime.Add(duration),
	}
}

type TokenMaker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
