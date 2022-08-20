package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
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
	row := r.db.QueryRowContext(ctx, query, args...)
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
	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return repoModels.Book{}, fmt.Errorf("Repository.GetBookById: to sql: %w", err)
	}
	var book repoModels.Book
	if ok := rows.Next(); !ok {
		return repoModels.Book{}, errors.New("no books")
	}
	err = rows.StructScan(&book)
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
	_, err = r.db.ExecContext(ctx, query, args...)

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
		return []repoModels.Book{}, fmt.Errorf("Repository.ListBooks: to sql: %w", err)
	}
	var books []repoModels.Book
	if err := r.db.SelectContext(ctx, &books, query, args...); err != nil {
		return nil, fmt.Errorf("Repository.ListBooks: to sql: %w", err)
	}
	return books, nil
}

func (r *repository) UpdateBook(ctx context.Context, book repoModels.Book) error {
	query, args, err := squirrel.Update("books").
		SetMap(
			map[string]interface{}{
				"title":     book.Title,
				"author_id": book.AuthorId,
			}).
		Where(
			squirrel.Eq{
				"id": book.Id,
			}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("Repository.UpdateBook: to sql")
	}
	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("Repository.UpdateBook: to sql")
	}
	return nil
}
