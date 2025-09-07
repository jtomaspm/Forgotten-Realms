package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Migrate(dbConfig Configuration, migrationsFolderPath string) (*Database, error) {
	ctx := context.Background()

	err := ensureDatabase(ctx, dbConfig)
	if err != nil {
		return nil, err
	}

	db := New(dbConfig.ConnectionString())

	err = ensureMigrationsTable(db.Pool, ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ensure migrations table: %w", err)
	}

	appliedMigrations, err := getAppliedMigrations(db.Pool, ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get applied migrations: %w", err)
	}

	migrationFiles, err := getMigrationFiles(migrationsFolderPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read migration files: %w", err)
	}

	for _, file := range migrationFiles {
		migrationName := strings.Replace(
			file, strings.Replace(migrationsFolderPath, "./", "", 1), "", 1)
		if _, alreadyApplied := appliedMigrations[migrationName]; alreadyApplied {
			log.Printf("Skipping already applied migration: %s\n", file)
			continue
		}

		err := applyMigration(db.Pool, ctx, file, migrationName)
		if err != nil {
			return nil, fmt.Errorf("failed to apply migration %s: %w", file, err)
		}

		log.Printf("Migration applied: %s\n", file)
	}

	log.Println("All migrations applied successfully.")
	return db, nil
}

func ensureDatabase(ctx context.Context, dbConfig Configuration) error {
	db := New(dbConfig.ConnectionStringWithDb("postgres"))
	defer db.Close()
	var exists bool
	err := db.Pool.QueryRow(ctx, `SELECT EXISTS (
		SELECT 1 FROM pg_database WHERE datname=$1
	)`, dbConfig.Database).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking if database exists: %w", err)
	}
	if exists {
		log.Println("Database already exists. Skipping creation.")
	} else {
		_, err = db.Pool.Exec(ctx, fmt.Sprintf("CREATE DATABASE %s", dbConfig.Database))
		if err != nil {
			return fmt.Errorf("error creating database: %w", err)
		}
		log.Printf("Database %s created successfully\n", dbConfig.Database)
	}
	return nil
}

func ensureMigrationsTable(pool *pgxpool.Pool, ctx context.Context) error {
	query := `
	CREATE TABLE IF NOT EXISTS migrations (
		id SERIAL PRIMARY KEY,
		name TEXT UNIQUE NOT NULL,
		applied_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := pool.Exec(ctx, query)
	return err
}

func getAppliedMigrations(pool *pgxpool.Pool, ctx context.Context) (map[string]bool, error) {
	rows, err := pool.Query(ctx, "SELECT name FROM migrations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	applied := make(map[string]bool)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		applied[name] = true
	}

	return applied, nil
}

func getMigrationFiles(migrationsFolderPath string) ([]string, error) {
	var migrationFiles []string

	err := filepath.WalkDir(migrationsFolderPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".sql") {
			migrationFiles = append(migrationFiles, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	sort.Strings(migrationFiles)

	return migrationFiles, nil
}

func applyMigration(pool *pgxpool.Pool, ctx context.Context, fileName, migrationName string) error {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	tx, err := pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, string(content))
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "INSERT INTO migrations (name) VALUES ($1)", migrationName)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
