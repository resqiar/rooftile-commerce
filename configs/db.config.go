package configs

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDBConfig() *pgxpool.Pool {
	DSN := os.Getenv("DSN")

	pool, err := pgxpool.New(context.Background(), DSN)
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	return pool
}
