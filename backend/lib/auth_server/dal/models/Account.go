package models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id         uuid.UUID
	ExternalId string
	Source     string
	Name       string
	Email      string
	Role       Role
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
