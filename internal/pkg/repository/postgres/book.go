package postgres

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	repoModels "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/models"
)

// TODO: один сервис или несколько???
// странное поведение, id инкрементиться при неудачной вставке, нужно решить это (можно делать проверку на существование автора перед этим)
func (r *repository) AddBook(ctx context.Context, book repoModels.Book) error {
	query, args, err := squirrel.Insert("books").
		Columns("title, author_id").
		Values(book.Title, book.AuthorId).
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
