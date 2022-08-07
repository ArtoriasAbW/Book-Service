package service

import (
	"context"
	"time"

	"github.com/pkg/errors"
	repoModels "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/models"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/service/models"
)

func (s *service) GetBook(ctx context.Context, id uint) (models.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*500))
	defer cancel()
	var err error
	book, err := s.Repository.GetBookById(ctx, id)
	return models.Book{
		Id:       book.Id,
		Title:    book.Title,
		AuthorId: book.AuthorId,
	}, err
}

func (s *service) AddBook(ctx context.Context, bookInput models.Book) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	var book repoModels.Book
	_, err := s.Repository.GetAuthorById(ctx, bookInput.AuthorId)
	if err != nil {
		return errors.Wrap(ErrValidation, "author with this id doesn't exist")
	}
	book.AuthorId = bookInput.AuthorId
	if bookInput.Title == "" {
		return errors.Wrap(ErrValidation, "field: [title] cannot be empty")
	}
	book.Title = bookInput.Title
	err = s.Repository.AddBook(ctx, book)
	return err
}

func (s *service) UpdateBook(ctx context.Context, book models.Book) error {
	return errors.New("not implemented")
}

func (s *service) DeleteBook(ctx context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	err := s.Repository.DeleteBook(ctx, id)
	return err
}

func (s *service) ListBooks(ctx context.Context) ([]models.Book, error) {
	return nil, errors.New("not implemented")
}
