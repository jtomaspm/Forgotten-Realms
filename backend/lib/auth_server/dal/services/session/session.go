package session

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Delete(ctx context.Context, pool *pgxpool.Pool, accountId uuid.UUID) error {
	query := `
		DELETE FROM sessions 
		WHERE account_id = $1;
	`
	_, err := pool.Exec(ctx, query, accountId)
	return err
}

func Create(ctx context.Context, pool *pgxpool.Pool, accountId uuid.UUID) (uuid.UUID, time.Time, error) {
	Delete(ctx, pool, accountId)
	var expires_at time.Time
	var token uuid.UUID
	query := `
		INSERT INTO sessions (account_id) 
		VALUES ($1) 
		RETURNING token, expires_at;
	`
	err := pool.QueryRow(ctx, query, accountId).Scan(&token, &expires_at)
	if err != nil {
		return token, expires_at, err
	}
	return token, expires_at, nil
}

func GetAccountId(ctx context.Context, pool *pgxpool.Pool, token uuid.UUID) (uuid.UUID, error) {
	var accountId uuid.UUID
	err := pool.QueryRow(ctx, `
		SELECT account_id
		FROM sessions
		WHERE token=$1
		AND expires_at > CURRENT_TIMESTAMP
		LIMIT 1	
	`, token).Scan(
		&accountId,
	)
	if err != nil {
		return accountId, err
	}
	return accountId, nil
}
