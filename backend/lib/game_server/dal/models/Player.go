package models

import (
	"time"

	"github.com/google/uuid"
)

type Player struct {
	Id        uuid.UUID
	Faction   Faction
	CreatedAt time.Time
	UpdatedAt time.Time
}
