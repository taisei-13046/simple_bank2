package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PaestoMaker struct {
	paesto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) < chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", chacha20poly1305.KeySize)
	}

	maker := &PaestoMaker{
		paesto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

func (maker *PaestoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	return maker.paesto.Encrypt(maker.symmetricKey, payload, nil)
}

func (maker *PaestoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := maker.paesto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
