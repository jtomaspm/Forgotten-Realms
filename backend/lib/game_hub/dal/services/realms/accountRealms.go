package realms

import (
	"backend/lib/game_hub/dal/models/queries"
	"backend/lib/game_hub/dal/models/views"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetByAccountId(ctx context.Context, pool *pgxpool.Pool, accountId uuid.UUID) ([]views.RegisteredRealm, error) {
	rows, err := pool.Query(ctx, `
		SELECT 
			r.id AS realm_id,
			r.name,
			r.api,
			(ar.account_id IS NOT NULL) AS registered,
			r.created_at
		FROM realms r
		LEFT JOIN account_realms ar
			ON r.id = ar.realm_id AND ar.account_id = $1
	`, accountId,
	)
	if err != nil {
		return []views.RegisteredRealm{}, err
	}
	realms := make([]views.RegisteredRealm, 0)
	for rows.Next() {
		var realm views.RegisteredRealm
		if rows.Scan(&realm.Id, &realm.Name, &realm.Api, &realm.Registered, &realm.CreatedAt) == nil {
			realms = append(realms, realm)
		}
	}
	return realms, nil
}

func RegisterAccount(ctx context.Context, pool *pgxpool.Pool, accountRealm *queries.CreateAccountRealm) error {
	_, err := pool.Exec(
		ctx,
		`
		INSERT INTO account_realms (account_id, realm_id) 
		VALUES ($1, $2) 
		`,
		accountRealm.AccountId, accountRealm.RealmId,
	)
	if err != nil {
		return err
	}
	return nil
}
