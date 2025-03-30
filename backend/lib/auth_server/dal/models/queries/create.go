package queries

import (
	"backend/lib/auth_server/dal/models"

	"github.com/google/uuid"
)

type CreateLogin struct {
	AccountId uuid.UUID
	IpAddress string
}

type CreateAccount struct {
	ExternalId             string
	Source                 string
	Name                   string
	Email                  string
	Role                   models.Role
	SendEmailNotifications bool
}

type CreateAccountProperties struct {
	AccountId              uuid.UUID
	SendEmailNotifications bool
}

type CreateSession struct {
	AccountId uuid.UUID
	Token     string
}
