package postgres

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	repoModels "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/models"
)

func (r *repository) AddAuthor(ctx context.Context, author repoModels.Author) error {
	query, args, err := squirrel.Insert("authors").
		Columns("author").
		Values(author.Name).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("Repository.AddAuthor: to sql: %w", err)
	}
	_, err = r.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("Repository.AddAuthor: to sql: %w", err)
	}
	return nil
}

func (r *repository) GetAuthorById(ctx context.Context, id uint) (repoModels.Author, error) {
	query, args, err := squirrel.Select("id, author").
		From("authors").
		Where(
			squirrel.Eq{
				"id": id,
			},
		).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return repoModels.Author{}, fmt.Errorf("Repository.GetAuthorById: to sql: %w", err)
	}

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return repoModels.Author{}, fmt.Errorf("Repository.GetAuthorById: to sql: %w", err)
	}
	var author repoModels.Author
	err = pgxscan.ScanOne(&author, rows)
	if err != nil {
		return repoModels.Author{}, fmt.Errorf("Repository.GetBookById: to sql: %w", err)
	}
	return author, nil
}
