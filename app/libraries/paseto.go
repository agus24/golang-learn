package libraries

import (
	"errors"
	"golang_gin/config"

	"strconv"
	"time"

	"github.com/o1egl/paseto"
)

type PasetoToken struct {
	paseto     *paseto.V2
	symmetric  []byte
	expiration time.Duration
}

func NewPasetoToken() *PasetoToken {
	expiration := config.PasetoExpirationTime
	return &PasetoToken{
		paseto:     paseto.NewV2(),
		symmetric:  []byte(config.PasetoSecret),
		expiration: expiration,
	}
}

func (t *PasetoToken) GenerateToken(userID int64) (string, error) {
	now := time.Now()
	claims := paseto.JSONToken{
		Subject:    strconv.FormatInt(userID, 10),
		IssuedAt:   now,
		Expiration: now.Add(t.expiration),
	}

	return t.paseto.Encrypt(t.symmetric, claims, "")
}

func (t *PasetoToken) VerifyToken(token string) (*paseto.JSONToken, error) {
	var newToken paseto.JSONToken
	var footer string

	err := t.paseto.Decrypt(token, t.symmetric, &newToken, &footer)
	if err != nil {
		return nil, err
	}

	if newToken.Expiration.Before(time.Now()) {
		return nil, err
	}

	return &newToken, nil
}

// ParseToken extracts and validates claims
func (p *PasetoToken) ParseToken(token string) (*paseto.JSONToken, error) {
	var claims paseto.JSONToken
	var footer string

	err := p.paseto.Decrypt(token, []byte(config.PasetoSecret), &claims, &footer)
	if err != nil {
		return nil, err
	}

	if time.Now().After(claims.Expiration) {
		return nil, errors.New("Token expired")
	}

	return &claims, nil
}
