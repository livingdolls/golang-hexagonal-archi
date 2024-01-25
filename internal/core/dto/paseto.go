package dto

import (
	"time"

	"github.com/google/uuid"
)

type TokenPayload struct {
	ID        uuid.UUID
	UserID    string
	Role      string
	IssuedAt  time.Time
	ExpiredAt time.Time
}
