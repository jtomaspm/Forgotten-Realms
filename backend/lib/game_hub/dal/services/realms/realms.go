package realms

import (
	"backend/lib/game_hub/dal/models"
	"backend/lib/game_hub/dal/models/queries"
	"backend/lib/game_hub/dal/models/views"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetAll(ctx context.Context, pool *pgxpool.Pool) ([]views.RegisteredRealm, error) {
	rows, err := pool.Query(ctx, `
		SELECT 
			id,
			name,
			status,
			FALSE AS registered,
			created_at
		FROM realms 
	`,
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

func GetById(ctx context.Context, pool *pgxpool.Pool, id uuid.UUID) (models.Realm, error) {
	var realm models.Realm
	var status string
	err := pool.QueryRow(ctx, `
		SELECT id, name, api, status, created_at, updated_at 
		FROM realms
		WHERE id=$1
		LIMIT 1	
	`, id,
	).Scan(
		&realm.Id,
		&realm.Name,
		&realm.Api,
		&status,
		&realm.CreatedAt,
		&realm.UpdatedAt,
	)
	if err != nil {
		return realm, err
	}
	realm.Status, err = models.FromString(status)
	if err != nil {
		return realm, err
	}
	return realm, nil
}

func GetByCreateQuery(ctx context.Context, pool *pgxpool.Pool, realm *queries.CreateRealm) (uuid.UUID, error) {
	var id uuid.UUID
	err := pool.QueryRow(ctx, `
		SELECT id
		FROM realms
		WHERE name=$1, api=$2
		LIMIT 1	
	`, realm.Name, realm.Api,
	).Scan(
		&id,
	)
	if err != nil {
		return id, err
	}
	return id, nil
}

func RegisterRealm(ctx context.Context, pool *pgxpool.Pool, realm *queries.CreateRealm) (uuid.UUID, error) {
	var id uuid.UUID
	err := pool.QueryRow(
		ctx,
		`
		INSERT INTO realms (name, api) 
		VALUES ($1, $2) 
		RETURNING id;
		`,
		realm.Name, realm.Api,
	).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}
