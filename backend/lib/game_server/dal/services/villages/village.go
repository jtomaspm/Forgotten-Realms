package villages

import (
	"backend/lib/game_server/configuration"
	"backend/pkg/database"
	"backend/pkg/sdk/game/enum"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Village struct {
	CoordX    int
	CoordY    int
	PlayerId  uuid.UUID
	Faction   enum.Faction
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetVillagesInRange(ctx context.Context, db database.Querier, minX, minY, maxX, maxY int) ([]Village, error) {
	if !(minX < configuration.MAP_SIZE || minX >= 0) {
		return []Village{}, fmt.Errorf("invalid minX")
	}
	if !(minY < configuration.MAP_SIZE || minY >= 0) {
		return []Village{}, fmt.Errorf("invalid minY")
	}
	if !(maxX < configuration.MAP_SIZE || maxX >= 0) {
		return []Village{}, fmt.Errorf("invalid maxX")
	}
	if !(maxY < configuration.MAP_SIZE || maxY >= 0) {
		return []Village{}, fmt.Errorf("invalid maxY")
	}
	rows, err := db.Query(ctx, `
		SELECT coord_x, coord_y, player_id, faction, created_at, updated_at
		FROM villages
		WHERE 
			coord_x >= $1 AND
			coord_y >= $2 AND
			coord_x <= $3 AND
			coord_y <= $4;
	`, minX, minY, maxX, maxY)
	if err != nil {
		return []Village{}, err
	}
	var villages []Village
	for rows.Next() {
		var faction string
		var village Village
		rows.Scan(&village.CoordX, &village.CoordY, &village.PlayerId, &faction, &village.CreatedAt, &village.UpdatedAt)
		village.Faction, err = enum.FactionFromString(faction)
		if err != nil {
			return []Village{}, err
		}
		villages = append(villages, village)
	}
	return villages, nil
}
