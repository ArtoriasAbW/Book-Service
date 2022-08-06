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
	User
	Review
}

type Book interface {
	GetBook(ctx context.Context, id uint) (models.Book, error)
	AddBook(ctx context.Context, bookInput models.Book) error
}

type Author interface {
	GetAuthor(ctx context.Context, id uint) (models.Author, error)
	AddAuthor(ctx context.Context, author models.Author) error
}

type User interface {
	GetUser(ctx context.Context, id uint) (models.User, error)
	AddUser(ctx context.Context, user models.User) error
}

type Review interface {
	GetReview(ctx context.Context, id uint) (models.Review, error)
	AddReview(ctx context.Context, review models.Review) (uint64, error)
	DeleteReview(ctx context.Context, id uint) error
	ListReviews(ctx context.Context, params models.ReviewListInput) ([]models.Review, error)
}

type service struct {
	Deps
}

type Deps struct {
	Repository repository.Interface
}
