package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	repoModels "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/models"
)

func (r *repository) AddAuthor(ctx context.Context, author repoModels.Author) (uint64, error) {
	query, args, err := squirrel.Insert("authors").
		Columns("author").
		Values(author.Name).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("Repository.AddAuthor: to sql: %w", err)
	}
	row := r.db.QueryRowContext(ctx, query, args...)
	var id uint64
	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("Repository.AddAuthor: to sql: %w", err)
	}
	return id, nil
}

func (r *repository) GetAuthorByID(ctx context.Context, id uint) (repoModels.Author, error) {
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

	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return repoModels.Author{}, fmt.Errorf("Repository.GetAuthorById: to sql: %w", err)
	}
	var author repoModels.Author
	ok := rows.Next()
	if !ok {
		return repoModels.Author{}, errors.New("no author")
	}
	err = rows.StructScan(&author)
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
	_, err = r.db.ExecContext(ctx, query, args...)
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
	if err := r.db.SelectContext(ctx, &authors, query, args...); err != nil {
		return nil, fmt.Errorf("Repository.ListAuthors: select: %w", err)
	}
	return authors, nil
}

func (r *repository) UpdateAuthor(ctx context.Context, author repoModels.Author) error {
	query, args, err := squirrel.Update("authors").
		SetMap(
			map[string]interface{}{
				"author": author.Name,
			}).
		Where(
			squirrel.Eq{
				"id": author.ID,
			}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("Repository.UpdateAuthor: to sql")
	}
	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("Repository.UpdateAuthor: to sql")
	}
	return nil
}
