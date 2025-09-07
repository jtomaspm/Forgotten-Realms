package players

import (
	"backend/lib/game_server/dal/models"
	"backend/lib/game_server/dal/models/queries"
	"backend/pkg/database"
	"backend/pkg/sdk/game/enum"
	"context"

	"github.com/google/uuid"
)

func Create(ctx context.Context, pool database.Querier, command *queries.CreatePlayer) (uuid.UUID, error) {
	var id uuid.UUID
	err := pool.QueryRow(
		ctx,
		`
		INSERT INTO players (id, faction) 
		VALUES ($1, $2) 
		RETURNING id;
		`,
		command.AccountId, command.Faction.String(),
	).Scan(&id)
	return id, err
}

func GetById(ctx context.Context, pool database.Querier, id uuid.UUID) (models.Player, error) {
	var player models.Player
	var faction string
	err := pool.QueryRow(ctx, `
		SELECT id, faction, created_at, updated_at 
		FROM players
		WHERE id=$1
		LIMIT 1	
	`, id,
	).Scan(
		&player.Id,
		&faction,
		&player.CreatedAt,
		&player.UpdatedAt,
	)
	if err != nil {
		return player, err
	}
	player.Faction, err = enum.FactionFromString(faction)
	if err != nil {
		return player, err
	}
	return player, nil
}
