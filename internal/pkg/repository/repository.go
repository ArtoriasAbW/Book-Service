package repository

import (
	"context"

	repoModels "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/models"
)

type Interface interface {
	Book
	// User
}

type Book interface {
	AddBook(ctx context.Context, b repoModels.Book) error
	GetBookById(ctx context.Context, id uint) (repoModels.Book, error)
	DeleteBook(ctx context.Context, id uint) error
}

type User interface {
	AddUser(ctx context.Context, user repoModels.User) error
	GetUserById(ctx context.Context, id uint) (repoModels.User, error)
	DeleteUserById(ctx context.Context, id uint) error
	ListUsers() ([]repoModels.User, error)
}

type Author interface {
	AddAuthor()
	GetAuthor()
	// DeleteAuthor()
	// ListAuthors()
}

type Review interface {
	AddReview()
	GetReview()
	DeleteReview()
	ListReviews()
}
