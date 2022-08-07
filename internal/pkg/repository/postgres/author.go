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
		return repoModels.Author{}, fmt.Errorf("Repository.GetAuthorById: to sql: %w", err)
	}
	return author, nil
}

func (r *repository) DeleteAuthor(ctx context.Context, id uint) error {
	query, args, err := squirrel.Delete("authors").
		Where(
			squirrel.Eq{
				"id": id,
			},
		).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("Repository.DeleteAuthor: to sql: %w", err)
	}
	_, err = r.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("Repository.DeleteAuthor: to sql: %w", err)
	}
	return nil
}

func (r *repository) ListAuthors(ctx context.Context, params repoModels.ListInput) ([]repoModels.Author, error) {
	preparedQuery := squirrel.Select("id", "author").
		From("authors").
		OrderBy("author " + params.Order).
		Offset(uint64(params.Offset))
	if params.Limit > 0 {
		preparedQuery = preparedQuery.Limit(uint64(params.Limit))
	}
	query, args, err := preparedQuery.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return []repoModels.Author{}, fmt.Errorf("Repository.ListAuthors: select: %w", err)
	}
	var authors []repoModels.Author
	if err := pgxscan.Select(ctx, r.pool, &authors, query, args...); err != nil {
		return nil, fmt.Errorf("Repository.ListAuthors: select: %w", err)
	}
	return authors, nil
}
