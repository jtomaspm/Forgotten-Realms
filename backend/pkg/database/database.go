package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Pool *pgxpool.Pool
}

type Configuration struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func (c *Configuration) ConnectionString() string {
	return "postgres://" + c.Username + ":" + c.Password + "@" + c.Host + ":" + c.Port + "/" + c.Database
}

func (c *Configuration) ConnectionStringWithDb(database string) string {
	return "postgres://" + c.Username + ":" + c.Password + "@" + c.Host + ":" + c.Port + "/" + database
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
