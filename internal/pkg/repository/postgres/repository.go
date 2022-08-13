package postgres

import "github.com/jackc/pgx/v4/pgxpool"

type repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *repository {
	return &repository{pool: pool}
}
