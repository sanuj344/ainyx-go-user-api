package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB(ctx context.Context, dbURL string) *pgxpool.Pool {
	db, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatal("unable to connect to database:", err)
	}

	if err := db.Ping(ctx); err != nil {
		log.Fatal("database ping failed:", err)
	}

	return db
}
