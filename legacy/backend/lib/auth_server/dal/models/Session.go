package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Id        uuid.UUID
	AccountId uuid.UUID
	Token     uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
