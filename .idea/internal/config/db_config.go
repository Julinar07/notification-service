package config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDBConnection() *pgxpool.Pool {
	connString := os.Getenv("postgres://postgres:postgres@localhost:5432/notification-service")

	if connString == "" {
		connString = "postgres://postgres:postgres@localhost:5432/postgres"
	}

	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v\n", err)
	}

	return pool
}