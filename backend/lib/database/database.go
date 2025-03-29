package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Pool *pgxpool.Pool
}

func New(connectionString string) *Database {
	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	return &Database{
		Pool: pool,
	}
}

func (db *Database) Close() {
	db.Pool.Close()
}
