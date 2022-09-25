package auth

import (
	"fmt"
	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto/v2"
	"time"
)

const symmetricKey = "TjWnZr4u7x!A%D*G-KaPdSgUkXp2s5v8"

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker() (TokenMaker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size must be exactly %d", chacha20poly1305.KeySize)
	}
	return &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}, nil
}

func (pasetoMaker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload := NewPayload(username, duration)

	return pasetoMaker.paseto.Encrypt(pasetoMaker.symmetricKey, payload, nil)
}

func (pasetoMaker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := pasetoMaker.paseto.Decrypt(token, pasetoMaker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
