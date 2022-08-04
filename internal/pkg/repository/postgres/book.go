package postgres

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/core/book/models"
)

func (r *Repository) AddBook(ctx context.Context, book models.Book) error {
	query, args, err := squirrel.Insert("books").
		Columns("title, author").
		Values(book.Title, book.Author).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("Repository.BookAdd: to sql: %w", err)
	}
	_, err = r.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("Repository.BookAdd: to sql: %w", err)
	}
	return nil
}

func (r *Repository) GetBookById(ctx context.Context, id uint) (models.Book, error) {
	query, args, err := squirrel.Select("id, title, author").
		From("books").
		Where(
			squirrel.Eq{
				"id": id,
			},
		).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return models.Book{}, fmt.Errorf("Repository.GetBookById: to sql: %w", err)
	}
	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return models.Book{}, fmt.Errorf("Repository.GetBookById: to sql: %w", err)
	}
	var book models.Book
	err = pgxscan.ScanOne(&book, rows)
	if err != nil {
		return models.Book{}, fmt.Errorf("Repository.GetBookById: to sql: %w", err)
	}
	return book, nil
}
