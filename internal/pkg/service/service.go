package service

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/service/models"
)

var ErrValidation = errors.New("invalid data")

func New(d Deps) *service {
	return &service{Deps: d}
}

type Interface interface {
	Book
	Author
}

type Book interface {
	GetBook(ctx context.Context, id uint) (models.Book, error)
	AddBook(ctx context.Context, bookInput models.Book) error
}

type Author interface {
	GetAuthor(ctx context.Context, id uint) (models.Author, error)
	AddAuthor(ctx context.Context, author models.Author) error
}

type service struct {
	Deps
}

type Deps struct {
	Repository repository.Interface
}
