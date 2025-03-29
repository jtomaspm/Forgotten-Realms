package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Id        uuid.UUID
	AccountId uuid.UUID
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
