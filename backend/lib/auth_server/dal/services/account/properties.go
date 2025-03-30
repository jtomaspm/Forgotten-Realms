package account

import (
	"backend/lib/auth_server/dal/models"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetProperties(ctx context.Context, pool *pgxpool.Pool, accountId uuid.UUID) (models.AccountProperties, error) {
	var properties models.AccountProperties
	err := pool.QueryRow(ctx, `
		SELECT account_id, verification_token, token_expires_at, email_verified, send_email_notifications, created_at, updated_at 
		FROM account_properties
		WHERE account_id=$1
		LIMIT 1	
	`, accountId).Scan(
		&properties.AccountId,
		&properties.VerificationToken,
		&properties.TokenExpiresAt,
		&properties.EmailVerified,
		&properties.SendEmailNotifications,
		&properties.CreatedAt,
		&properties.UpdatedAt,
	)
	if err != nil {
		return properties, err
	}
	return properties, nil
}

func CreateProperties(ctx context.Context, pool *pgxpool.Pool, accountId uuid.UUID) (uuid.UUID, error) {
	var verificationToken uuid.UUID
	query := `
		INSERT INTO account_properties (account_id) 
		VALUES ($1) 
		RETURNING verification_token;
	`
	err := pool.QueryRow(
		ctx, query, accountId,
	).Scan(&verificationToken)
	if err != nil {
		return verificationToken, err
	}
	return verificationToken, nil
}

func Verify(ctx context.Context, pool *pgxpool.Pool, accountId uuid.UUID, token uuid.UUID) error {
	properties, err := GetProperties(ctx, pool, accountId)
	if err != nil {
		return err
	}
	if properties.EmailVerified {
		return fmt.Errorf("email already verified")
	}
	if properties.TokenExpiresAt.Before(time.Now()) {
		return fmt.Errorf("token expired")
	}
	if properties.VerificationToken != token {
		return fmt.Errorf("invalid token")
	}

	query := `
		UPDATE account_properties 
		SET email_verified=TRUE
		WHERE account_id=$1;
	`
	_, err = pool.Exec(
		ctx, query, accountId,
	)
	return err
}
