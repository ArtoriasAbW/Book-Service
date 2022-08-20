package postgres

import (
	"github.com/jmoiron/sqlx"
	repoPkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) repoPkg.Interface {
	return &repository{db: db}
}
