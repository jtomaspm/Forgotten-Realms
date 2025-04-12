package queries

import (
	"backend/lib/game_server/dal/models"

	"github.com/google/uuid"
)

type CreatePlayer struct {
	AccountId uuid.UUID
	Faction   models.Faction
}
