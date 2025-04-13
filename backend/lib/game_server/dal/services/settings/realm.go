package settings

import (
	"backend/pkg/database"
	"context"
)

type Realm struct {
	Speed     float32 `json:"speed"`
	UnitSpeed float32 `json:"unit_speed"`
}

func (l *Realm) Sync(ctx context.Context, pool database.Querier) error {
	_, err := pool.Exec(
		ctx,
		`
		INSERT INTO settings_realm (speed, unit_speed)
		VALUES ($1, $2)
		ON CONFLICT (created_at) DO UPDATE SET
			speed = EXCLUDED.speed,
			unit_speed = EXCLUDED.unit_speed,
			updated_at = CURRENT_TIMESTAMP;
		`,
		l.Speed, l.UnitSpeed,
	)
	return err
}
