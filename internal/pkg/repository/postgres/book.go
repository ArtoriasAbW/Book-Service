package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	repoModels "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/models"
)

func (r *repository) AddBook(ctx context.Context, book repoModels.Book) (uint64, error) {
	query, args, err := squirrel.Insert("books").
		Columns("title, author_id").
		Values(book.Title, book.AuthorId).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("Repository.BookAdd: to sql: %w", err)
	}
	row := r.pool.QueryRow(ctx, query, args...)
	var id uint64
	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("Repository.BookAdd: to sql: %w", err)
	}
	return id, nil
}

func (r *repository) GetBookById(ctx context.Context, id uint) (repoModels.Book, error) {
	query, args, err := squirrel.Select("id, title, author_id").
		From("books").
		Where(
			squirrel.Eq{
				"id": id,
			},
		).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return repoModels.Book{}, fmt.Errorf("Repository.GetBookById: to sql: %w", err)
	}
	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return repoModels.Book{}, fmt.Errorf("Repository.GetBookById: to sql: %w", err)
	}
	var book repoModels.Book
	err = pgxscan.ScanOne(&book, rows)
	if err != nil {
		return repoModels.Book{}, fmt.Errorf("Repository.GetBookById: to sql: %w", err)
	}
	return book, nil
}

func (r *repository) DeleteBook(ctx context.Context, id uint) error {
	query, args, err := squirrel.Delete("books").
		Where(
			squirrel.Eq{
				"id": id,
			},
		).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("Repository.DeleteBook: to sql: %w", err)
	}
	_, err = r.pool.Exec(ctx, query, args...)

	if err != nil {
		return fmt.Errorf("Repository.DeleteBook: to sql: %w", err)
	}
	return nil
}

func (r *repository) ListBooks(ctx context.Context, params repoModels.ListInput) ([]repoModels.Book, error) {
	preparedQuery := squirrel.Select("id", "title", "author_id").
		From("books").
		OrderBy("title " + params.Order).
		Offset(uint64(params.Offset))
	if params.Limit > 0 {
		preparedQuery = preparedQuery.Limit(uint64(params.Limit))
	}
	query, args, err := preparedQuery.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return []repoModels.Book{}, fmt.Errorf("Repository.ListBooks: select: %w", err)
	}
	var books []repoModels.Book
	if err := pgxscan.Select(ctx, r.pool, &books, query, args...); err != nil {
		return nil, fmt.Errorf("Repository.ListBooks: select: %w", err)
	}
	return books, nil
}

func (r *repository) UpdateBook(ctx context.Context, book repoModels.Book) error {
	return errors.New("not implemented")
}
