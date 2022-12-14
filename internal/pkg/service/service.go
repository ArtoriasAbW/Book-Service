package service

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/service/models"
)

var ErrValidation = errors.New("invalid data")

func New(d Deps) *Implementation {
	return &Implementation{Deps: d}
}

type Interface interface {
	book
	author
	user
	review
}

type book interface {
	GetBook(ctx context.Context, id uint) (models.Book, error)
	AddBook(ctx context.Context, bookInput models.Book) (uint64, error)
	DeleteBook(ctx context.Context, id uint) error
	ListBooks(ctx context.Context, params models.ListInput) ([]models.Book, error)
	UpdateBook(ctx context.Context, bookInput models.Book) error
}

type author interface {
	GetAuthor(ctx context.Context, id uint) (models.Author, error)
	AddAuthor(ctx context.Context, author models.Author) (uint64, error)
	DeleteAuthor(ctx context.Context, id uint) error
	ListAuthors(ctx context.Context, params models.ListInput) ([]models.Author, error)
	UpdateAuthor(ctx context.Context, authorInput models.Author) error
}

type user interface {
	GetUser(ctx context.Context, id uint) (models.User, error)
	AddUser(ctx context.Context, user models.User) (uint64, error)
	DeleteUser(ctx context.Context, id uint) error
	ListUsers(ctx context.Context, params models.ListInput) ([]models.User, error)
	UpdateUser(ctx context.Context, userInput models.User) error
}

type review interface {
	GetReview(ctx context.Context, id uint) (models.Review, error)
	AddReview(ctx context.Context, review models.Review) (uint64, error)
	DeleteReview(ctx context.Context, id uint) error
	ListReviews(ctx context.Context, params models.ListInput) ([]models.Review, error)
	UpdateReview(ctx context.Context, reviewInput models.Review) error
}

type Implementation struct {
	Deps
}

type Deps struct {
	Repository repository.Interface
}
