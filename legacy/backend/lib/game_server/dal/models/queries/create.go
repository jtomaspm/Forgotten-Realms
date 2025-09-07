package queries

import (
	"backend/pkg/sdk/game/enum"

	"github.com/google/uuid"
)

type CreatePlayer struct {
	AccountId uuid.UUID
	Faction   enum.Faction
}
