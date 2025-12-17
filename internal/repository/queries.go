package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sanuj344/ainyx-go-user-api/db/sqlc"
)

type Repository struct {
	Queries *sqlc.Queries
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Queries: sqlc.New(db),
	}
}
