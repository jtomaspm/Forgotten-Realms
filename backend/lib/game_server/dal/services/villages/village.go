package villages

import (
	"backend/lib/game_server/dal/services/settings"
	"backend/pkg/database"
	"backend/pkg/sdk/game/enum"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Coord struct {
	CoordX int `json:"coord_x"`
	CoordY int `json:"coord_y"`
}

type NewVillage struct {
	CoordX   int
	CoordY   int
	PlayerId uuid.UUID
}

type Village struct {
	CoordX    int
	CoordY    int
	PlayerId  uuid.UUID
	Faction   enum.Faction
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateVillage(ctx context.Context, db database.Querier, village NewVillage) error {
	if village.CoordX < 0 || village.CoordY < 0 {
		return fmt.Errorf("invalid coordinates: (%d, %d)", village.CoordX, village.CoordY)
	}
	if village.PlayerId == uuid.Nil {
		return fmt.Errorf("player ID cannot be nil")
	}

	_, err := db.Exec(ctx, `
		INSERT INTO villages (coord_x, coord_y, player_id)
		VALUES ($1, $2, $3);
	`, village.CoordX, village.CoordY, village.PlayerId)
	return err
}

func GetVillagesInRange(ctx context.Context, db database.Querier, minX, minY, maxX, maxY int, rs settings.Realm) ([]Village, error) {
	if !(minX < rs.MapSize || minX >= 0) {
		return []Village{}, fmt.Errorf("invalid minX")
	}
	if !(minY < rs.MapSize || minY >= 0) {
		return []Village{}, fmt.Errorf("invalid minY")
	}
	if !(maxX < rs.MapSize || maxX >= 0) {
		return []Village{}, fmt.Errorf("invalid maxX")
	}
	if !(maxY < rs.MapSize || maxY >= 0) {
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
