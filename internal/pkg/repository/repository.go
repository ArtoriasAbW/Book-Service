//go:generate mockgen ./repository.go -destination ./mocks/repository.go -package=mock_repository
package repository

import (
	"context"

	repoModels "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/models"
)

type Interface interface {
	book
	author
	user
	review
}

type book interface {
	AddBook(ctx context.Context, b repoModels.Book) (uint64, error)
	GetBookByID(ctx context.Context, id uint) (repoModels.Book, error)
	DeleteBook(ctx context.Context, id uint) error
	ListBooks(ctx context.Context, params repoModels.ListInput) ([]repoModels.Book, error)
	UpdateBook(ctx context.Context, book repoModels.Book) error
}

type user interface {
	AddUser(ctx context.Context, user repoModels.User) (uint64, error)
	GetUserByID(ctx context.Context, id uint) (repoModels.User, error)
	DeleteUser(ctx context.Context, id uint) error
	ListUsers(ctx context.Context, params repoModels.ListInput) ([]repoModels.User, error)
	UpdateUser(ctx context.Context, newUser repoModels.User) error
}

type author interface {
	AddAuthor(ctx context.Context, author repoModels.Author) (uint64, error)
	GetAuthorByID(ctx context.Context, id uint) (repoModels.Author, error)
	DeleteAuthor(ctx context.Context, id uint) error
	ListAuthors(ctx context.Context, params repoModels.ListInput) ([]repoModels.Author, error)
	UpdateAuthor(ctx context.Context, author repoModels.Author) error
}

type review interface {
	AddReview(ctx context.Context, review repoModels.Review) (uint64, error)
	GetReviewByID(ctx context.Context, id uint) (repoModels.Review, error)
	DeleteReview(ctx context.Context, id uint) error
	ListReviews(ctx context.Context, params repoModels.ListInput) ([]repoModels.Review, error)
	UpdateReview(ctx context.Context, review repoModels.Review) error
}
