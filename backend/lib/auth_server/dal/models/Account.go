package models

import (
	"backend/pkg/core/models"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id         uuid.UUID
	ExternalId string
	Source     string
	Name       string
	Email      string
	Role       models.Role
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
