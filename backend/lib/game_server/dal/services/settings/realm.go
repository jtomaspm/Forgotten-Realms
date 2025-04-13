package settings

import (
	"backend/pkg/database"
	"context"
	"fmt"
)

type Realm struct {
	Speed     float32 `json:"speed"`
	UnitSpeed float32 `json:"unit_speed"`
}

func (l *Realm) Sync(ctx context.Context, pool database.Querier) error {
	var count int
	err := pool.QueryRow(ctx, `SELECT COUNT(*) FROM settings_realm`).Scan(&count)
	if err != nil {
		return fmt.Errorf("counting settings_realm: %w", err)
	}

	if count == 0 {
		_, err = pool.Exec(ctx, `
			INSERT INTO settings_realm (speed, unit_speed)
			VALUES ($1, $2)
		`, l.Speed, l.UnitSpeed)
		if err != nil {
			return fmt.Errorf("inserting settings_realm: %w", err)
		}
	} else {
		_, err = pool.Exec(ctx, `
			UPDATE settings_realm
			SET speed = $1,
				unit_speed = $2,
				updated_at = CURRENT_TIMESTAMP
		`, l.Speed, l.UnitSpeed)
		if err != nil {
			return fmt.Errorf("updating settings_realm: %w", err)
		}
	}

	return nil
}
