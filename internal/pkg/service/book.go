package service

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	repoModels "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/models"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/service/models"
)

func (c *service) GetBook(ctx context.Context, id uint) (models.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*50))
	defer cancel()
	var err error
	book, err := c.Repository.GetBookById(ctx, id)
	return models.Book{
		Id:       book.Id,
		Title:    book.Title,
		AuthorId: book.AuthorId,
	}, err
}

func (c *service) AddBook(ctx context.Context, bookInput models.Book) error {
	fmt.Println("service:", bookInput)
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	var book repoModels.Book
	_, err := c.Repository.GetAuthorById(ctx, bookInput.AuthorId)
	if err != nil {
		return errors.Wrap(ErrValidation, "author with this id doesn't exist")
	}
	book.AuthorId = bookInput.AuthorId
	if bookInput.Title == "" {
		return errors.Wrap(ErrValidation, "field: [title] cannot be empty")
	}
	book.Title = bookInput.Title
	err = c.Repository.AddBook(ctx, book)
	return err
}

func (c *service) UpdateBook(ctx context.Context, book models.Book) error {
	return errors.New("not implemented")
}

func (c *service) DeleteBook(ctx context.Context, id uint) error {
	return errors.New("not implemented")

}

func (c *service) ListBooks(ctx context.Context) ([]models.Book, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*50))
	// defer cancel()
	// var books []models.Book
	// c.repository.
	return nil, errors.New("not implemented")
}
