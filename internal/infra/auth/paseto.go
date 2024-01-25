package auth

import (
	"errors"
	"fmt"
	"gotest/internal/core/dto"
	"gotest/internal/core/port/service"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

type PasetoToken struct {
	paseto      *paseto.V2
	symetricKey []byte
	duration    time.Duration
}

// ValidateToken implements service.TokenService.
func (p *PasetoToken) ValidateToken(token string) (*dto.TokenPayload, error) {
	var payload dto.TokenPayload

	err := p.paseto.Decrypt(token, p.symetricKey, &payload, nil)

	if err != nil {
		return nil, errors.New("invalid token")
	}

	isExperied := time.Now().After(payload.ExpiredAt)

	if isExperied {
		return nil, errors.New("token experied")
	}

	return &payload, err
}

// CreateToken implements service.TokenService.
func (p *PasetoToken) CreateToken(person *dto.PersonDTO) (string, error) {
	newid, err := uuid.NewRandom()

	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("error uuid creation")
	}

	payload := dto.TokenPayload{
		ID:        newid,
		UserID:    person.PersonsID,
		Role:      "admin",
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(p.duration),
	}

	token, err := p.paseto.Encrypt(p.symetricKey, payload, nil)

	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("error encrypt creation")
	}

	return token, nil
}

func New() (service.TokenService, error) {
	symetricKey := "12345678901234567890123456789012"
	durationStr := "15m"

	fmt.Println("sini")

	validSymetricKey := len(symetricKey) == chacha20poly1305.KeySize

	if !validSymetricKey {
		return nil, errors.New("Invalid token key size")
	}

	duration, err := time.ParseDuration(durationStr)

	if err != nil {
		return nil, err
	}

	return &PasetoToken{
		paseto.NewV2(),
		[]byte(symetricKey),
		duration,
	}, nil
}
