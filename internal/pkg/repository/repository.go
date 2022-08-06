package repository

import (
	"context"

	repoModels "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/models"
)

type Interface interface {
	Book
	Author
	User
	Review
}

type Book interface {
	AddBook(ctx context.Context, b repoModels.Book) error
	GetBookById(ctx context.Context, id uint) (repoModels.Book, error)
	DeleteBook(ctx context.Context, id uint) error
}

type User interface {
	AddUser(ctx context.Context, user repoModels.User) error
	GetUserById(ctx context.Context, id uint) (repoModels.User, error)
	// DeleteUserById(ctx context.Context, id uint) error
	// ListUsers() ([]repoModels.User, error)
}

type Author interface {
	AddAuthor(ctx context.Context, author repoModels.Author) error
	GetAuthorById(ctx context.Context, id uint) (repoModels.Author, error)
	// DeleteAuthor()
	// ListAuthors()
}

type Review interface {
	AddReview(ctx context.Context, review repoModels.Review) (uint64, error)
	GetReviewById(ctx context.Context, id uint) (repoModels.Review, error)
	DeleteReview(ctx context.Context, id uint) error
	ListReviews(ctx context.Context, params repoModels.ReviewListInput) ([]repoModels.Review, error)
}
