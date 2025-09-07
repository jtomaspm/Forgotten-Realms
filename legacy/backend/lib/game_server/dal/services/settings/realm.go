package settings

import (
	"backend/pkg/database"
	"context"
	"fmt"
)

type Realm struct {
	Speed            float32 `json:"speed"`
	UnitSpeed        float32 `json:"unit_speed"`
	ChunkSize        int     `json:"chunk_size"`
	ChunkFillPercent int     `json:"chunk_fill_percent"`
	MapSize          int     `json:"map_size"`
}

func GetRealmSettings(ctx context.Context, db database.Querier) (Realm, error) {
	var realm Realm
	err := db.QueryRow(ctx, `
			SELECT speed, unit_speed, chunk_size, chunk_fill_percent, map_size
	 		FROM settings_realm
			LIMIT 1
		`).Scan(&realm.Speed, &realm.UnitSpeed, &realm.ChunkSize, &realm.ChunkFillPercent, &realm.MapSize)
	if err != nil {
		return Realm{}, fmt.Errorf("get settings_realm: %w", err)
	}
	return realm, nil
}

func (l *Realm) Sync(ctx context.Context, db database.Querier) error {
	var count int
	err := db.QueryRow(ctx, `SELECT COUNT(*) FROM settings_realm`).Scan(&count)
	if err != nil {
		return fmt.Errorf("counting settings_realm: %w", err)
	}

	if count == 0 {
		_, err = db.Exec(ctx, `
			INSERT INTO settings_realm (speed, unit_speed, chunk_size, chunk_fill_percent, map_size)
			VALUES ($1, $2, $3, $4, $5)
		`, l.Speed, l.UnitSpeed, l.ChunkSize, l.ChunkFillPercent, l.MapSize)
		if err != nil {
			return fmt.Errorf("inserting settings_realm: %w", err)
		}
	} else {
		_, err = db.Exec(ctx, `
			UPDATE settings_realm
			SET speed = $1,
				unit_speed = $2,
				chunk_size = $3,
				chunk_fill_percent = $4,
				map_size = $5,
				updated_at = CURRENT_TIMESTAMP
		`, l.Speed, l.UnitSpeed, l.ChunkSize, l.ChunkFillPercent, l.MapSize)
		if err != nil {
			return fmt.Errorf("updating settings_realm: %w", err)
		}
	}
	return nil
}
