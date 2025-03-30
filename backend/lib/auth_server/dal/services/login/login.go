package login

import (
	"backend/lib/auth_server/dal/models/queries"
	"backend/lib/auth_server/dal/services/account"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Create(ctx context.Context, pool *pgxpool.Pool, login *queries.CreateLogin) (uuid.UUID, error) {
	var id uuid.UUID

	properties, err := account.GetProperties(ctx, pool, login.AccountId)
	if err != nil {
		return id, err
	}
	if !properties.EmailVerified {
		return id, fmt.Errorf("email not verified")
	}

	query := `
		INSERT INTO logins (account_id, ip_address) 
		VALUES ($1, $2) 
		RETURNING id;
	`
	err = pool.QueryRow(
		ctx, query, login.AccountId, login.IpAddress,
	).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}
