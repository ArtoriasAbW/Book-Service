package postgres

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jmoiron/sqlx"
	repoPkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository"
)

type repository struct {
	pool *pgxpool.Pool
	db   *sqlx.DB
}

func NewRepository(pool *pgxpool.Pool, db *sqlx.DB) repoPkg.Interface {
	return &repository{pool: pool, db: db}
}
