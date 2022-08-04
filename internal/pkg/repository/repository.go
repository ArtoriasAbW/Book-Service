package repository

import (
	"context"

	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/core/book/models"
)

type Interface interface {
	Book
}

type Book interface {
	AddBook(ctx context.Context, b models.Book) error
	GetBookById(ctx context.Context, id uint) (models.Book, error)
}
