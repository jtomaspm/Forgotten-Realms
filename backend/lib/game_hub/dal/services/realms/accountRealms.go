package realms

import (
	"backend/lib/game_hub/dal/models/views"
	"backend/pkg/sdk/hub/realms"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetPlayableRealms(ctx context.Context, pool *pgxpool.Pool, accountId uuid.UUID) ([]views.PlayableRealm, error) {
	rows, err := pool.Query(ctx, `
		SELECT 
			r.id AS realm_id,
			r.name,
			r.api
		FROM realms r
		JOIN account_realms ar
			ON r.id = ar.realm_id AND ar.account_id = $1
		WHERE r.status!='ended'
	`, accountId,
	)
	if err != nil {
		return []views.PlayableRealm{}, err
	}
	realms := make([]views.PlayableRealm, 0)
	for rows.Next() {
		var realm views.PlayableRealm
		if rows.Scan(&realm.Id, &realm.Name, &realm.Api) == nil {
			realms = append(realms, realm)
		}
	}
	return realms, nil
}

func GetAllByAccountId(ctx context.Context, pool *pgxpool.Pool, accountId uuid.UUID) ([]views.RegisteredRealm, error) {
	rows, err := pool.Query(ctx, `
		SELECT 
			r.id AS realm_id,
			r.name,
			r.status,
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
		if rows.Scan(&realm.Id, &realm.Name, &realm.Status, &realm.Registered, &realm.CreatedAt) == nil {
			realms = append(realms, realm)
		}
	}
	return realms, nil
}

func RegisterAccount(ctx context.Context, pool *pgxpool.Pool, accountRealm *realms.RegisterAccountRequestBody) error {
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
