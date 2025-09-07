package models

import (
	"backend/pkg/sdk/game/enum"
	"time"

	"github.com/google/uuid"
)

type Player struct {
	Id        uuid.UUID
	Faction   enum.Faction
	CreatedAt time.Time
	UpdatedAt time.Time
}
