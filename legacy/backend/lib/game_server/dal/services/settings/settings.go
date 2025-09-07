package settings

import (
	"backend/pkg/database"
	"context"
)

type Setting interface {
	Sync(ctx context.Context, pool database.Querier) error
}
