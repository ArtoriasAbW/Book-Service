package service

import (
	"context"
	"strings"
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

func (s *service) AddBook(ctx context.Context, bookInput models.Book) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	var book repoModels.Book
	_, err := s.Repository.GetAuthorById(ctx, bookInput.AuthorId)
	if err != nil {
		return 0, errors.Wrap(ErrValidation, "author with this id doesn't exist")
	}
	book.AuthorId = bookInput.AuthorId
	if bookInput.Title == "" {
		return 0, errors.Wrap(ErrValidation, "field: [title] cannot be empty")
	}
	book.Title = bookInput.Title
	id, err := s.Repository.AddBook(ctx, book)
	return id, err
}

func (s *service) UpdateBook(ctx context.Context, bookInput models.Book) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	var book repoModels.Book
	_, err := s.Repository.GetAuthorById(ctx, bookInput.AuthorId)
	if err != nil {
		return errors.Wrap(ErrValidation, "author with this id doesn't exist")
	}
	book.Id = bookInput.Id
	_, err = s.Repository.GetBookById(ctx, book.Id)
	if err != nil {
		return errors.Wrap(ErrValidation, "book with this id doesn't exist")
	}
	book.AuthorId = bookInput.AuthorId
	if bookInput.Title == "" {
		return errors.Wrap(ErrValidation, "field: [title] cannot be empty")
	}
	book.Title = bookInput.Title
	err = s.Repository.UpdateBook(ctx, book)
	return err
}

func (s *service) DeleteBook(ctx context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	err := s.Repository.DeleteBook(ctx, id)
	return err
}

func (s *service) ListBooks(ctx context.Context, params models.ListInput) ([]models.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	var repoParams repoModels.ListInput
	repoParams.Offset = params.Offset
	repoParams.Limit = params.Limit
	if strings.ToLower(params.Order) == "desc" {
		repoParams.Order = "DESC"
	} else {
		repoParams.Order = "ASC"
	}
	booksRepo, err := s.Repository.ListBooks(ctx, repoParams)
	if err != nil {
		return nil, err
	}
	books := make([]models.Book, len(booksRepo))
	for i, book := range booksRepo {
		books[i] = models.Book{
			Id:       book.Id,
			Title:    book.Title,
			AuthorId: book.AuthorId,
		}
	}
	return books, nil
}
