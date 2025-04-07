package models

import (
	"time"

	"github.com/google/uuid"
)

type Player struct {
	Id        uuid.UUID
	AccountId uuid.UUID
	Faction   FactionEnum
	CreatedAt time.Time
	UpdatedAt time.Time
}
