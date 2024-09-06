package config

import (
    "context"
    "log"
    "os"

    "github.com/jackc/pgx/v4/pgxpool"
)

var Pool *pgxpool.Pool

// InitPgxPool initializes the PostgreSQL connection. It parses the connection string and creates a
// connection pool using the pgxpool package.
func InitPgxPool() {
    ctx := context.Background()
    dbURL := os.Getenv("DB_URL")

    if dbURL == "" {
        log.Fatal("DB_URL environment variable not set")
    }

    config, err := pgxpool.ParseConfig(dbURL)
    if err != nil {
        log.Fatalf("Unable to parse connection string: %v", err)
    }

    Pool, err = pgxpool.ConnectConfig(ctx, config)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v", err)
    }

    log.Println("Successfully connected to the database with pgxpool!")
}
