package models

import (
	"time"

	"github.com/google/uuid"
)

type AccountRealm struct {
	AccountId uuid.UUID
	RealmId   uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
