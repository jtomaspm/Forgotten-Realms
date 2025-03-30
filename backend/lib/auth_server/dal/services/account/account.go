package account

import (
	"backend/lib/auth_server/dal/models"
	"backend/lib/auth_server/dal/models/queries"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetById(ctx context.Context, pool *pgxpool.Pool, id uuid.UUID) (models.Account, error) {
	var account models.Account
	var role string
	err := pool.QueryRow(ctx, `
		SELECT id, external_id, source, name, email, role, created_at, updated_at 
		FROM accounts
		WHERE id=$1
		LIMIT 1	
	`, id,
	).Scan(
		&account.Id,
		&account.ExternalId,
		&account.Source,
		&account.Name,
		&account.Email,
		&role,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err != nil {
		return account, err
	}
	account.Role, err = models.FromString(role)
	if err != nil {
		return account, err
	}
	return account, nil
}

func GetByExternalId(ctx context.Context, pool *pgxpool.Pool, account *queries.GetAccountByExternalId) (models.Account, error) {
	var result models.Account
	var role string
	err := pool.QueryRow(ctx, `
		SELECT id, external_id, source, name, email, role, created_at, updated_at 
		FROM accounts
		WHERE external_id=$1 AND source=$2
		LIMIT 1	
	`, account.ExternalId, account.Source,
	).Scan(
		&result.Id,
		&result.ExternalId,
		&result.Source,
		&result.Name,
		&result.Email,
		&role,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if err != nil {
		return result, err
	}
	result.Role, err = models.FromString(role)
	if err != nil {
		return result, err
	}
	return result, nil
}

func Create(ctx context.Context, pool *pgxpool.Pool, account *queries.CreateAccount) (uuid.UUID, error) {
	var id uuid.UUID
	err := pool.QueryRow(
		ctx,
		`
		INSERT INTO accounts (external_id, source, name, email, role) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id;
		`,
		account.ExternalId, account.Source, account.Name, account.Email, account.Role.String(),
	).Scan(&id)
	if err != nil {
		return id, err
	}
	return CreateProperties(ctx, pool, &queries.CreateAccountProperties{
		AccountId:              id,
		SendEmailNotifications: account.SendEmailNotifications,
	})
}
