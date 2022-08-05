package service

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	repoModels "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/models"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/service/models"
)

var ErrValidation = errors.New("invalid data")

func (c *service) GetBook(id uint) (models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*50))
	defer cancel()
	var err error
	book, err := c.Repository.GetBookById(ctx, id)
	return models.Book{
		Id:       book.Id,
		Title:    book.Title,
		AuthorId: book.AuthorId,
	}, err
}

func (c *service) AddBook(bookInput models.Book) error {
	fmt.Println("service:", bookInput)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*1000))
	defer cancel()
	var book repoModels.Book
	// TODO: check author id
	book.AuthorId = bookInput.AuthorId
	if bookInput.Title == "" {
		return errors.Wrap(ErrValidation, "field: [title] cannot be empty")
	}
	book.Title = bookInput.Title
	err := c.Repository.AddBook(ctx, book)
	return err
}

func (c *service) Update(book models.Book) error {
	return errors.New("not implemented")
}

func (c *service) Delete(id uint) error {
	return errors.New("not implemented")

}

func (c *service) List() ([]models.Book, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*50))
	// defer cancel()
	// var books []models.Book
	// c.repository.
	return nil, errors.New("not implemented")
}
