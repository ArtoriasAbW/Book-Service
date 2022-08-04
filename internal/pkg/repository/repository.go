package repository

import (
	"context"

	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/core/book/models"
)

type Interface interface {
	Book
	User
	Author
}

type Book interface {
	AddBook(ctx context.Context, b models.Book) error
	GetBookById(ctx context.Context, id uint) (models.Book, error)
}

type User interface {
	AddUser()
	GetUser()
	DeleteUser()
	ListUsers()
}

type Author interface {
	AddAuthor()
	GetAuthor()
	DeleteAuthor()
	ListAuthors()
}

type Review interface {
	AddReview()
	GetReview()
	DeleteReview()
	ListReviews()
}
