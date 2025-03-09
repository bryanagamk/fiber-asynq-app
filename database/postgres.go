package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	dsn := "postgres://postgres:postgres@localhost:5432/fiber_asynq"
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Unable to parse DB config: %v\n", err)
	}

	config.MaxConns = 50
	config.MinConns = 10

	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to DB: %v\n", err)
	}

	log.Println("Connected to DB")
}
