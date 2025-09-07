package models

import (
	"time"

	"github.com/google/uuid"
)

type AccountProperties struct {
	AccountId              uuid.UUID
	VerificationToken      uuid.UUID
	TokenExpiresAt         time.Time
	EmailVerified          bool
	SendEmailNotifications bool
	CreatedAt              time.Time
	UpdatedAt              time.Time
}
