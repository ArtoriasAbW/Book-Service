package service

import (
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/service/models"
	pb "gitlab.ozon.dev/ArtoriasAbW/homework-01/pkg/api"
)

func (s *Implementation) GetBook(ctx context.Context, id uint) (models.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*500))
	defer cancel()
	var err error
	bookGetResponce, err := s.Repository.BookGet(ctx, &pb.BookGetRequest{
		Id: uint64(id),
	})
	if err != nil {
		return models.Book{}, err
	}
	return models.Book{
		AuthorId: uint(bookGetResponce.GetAuthorId()),
		Title:    bookGetResponce.GetTitle(),
	}, nil
}

func (s *Implementation) AddBook(ctx context.Context, bookInput models.Book) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	_, err := s.Repository.AuthorGet(ctx, &pb.AuthorGetRequest{
		Id: uint64(bookInput.AuthorId),
	})
	if err != nil {
		return 0, errors.Wrap(ErrValidation, "author with this id doesn't exist")
	}
	if bookInput.Title == "" {
		return 0, errors.Wrap(ErrValidation, "field: [title] cannot be empty")
	}
	bookResponce, err := s.Repository.BookCreate(ctx, &pb.BookCreateRequest{
		Title:    bookInput.Title,
		AuthorId: uint64(bookInput.AuthorId),
	})
	if err != nil {
		return 0, err
	}
	return bookResponce.GetId(), nil
}

func (s *Implementation) UpdateBook(ctx context.Context, bookInput models.Book) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	_, err := s.Repository.AuthorGet(ctx, &pb.AuthorGetRequest{
		Id: uint64(bookInput.AuthorId),
	})
	if err != nil {
		return errors.Wrap(ErrValidation, "author with this id doesn't exist")
	}
	_, err = s.Repository.BookGet(ctx, &pb.BookGetRequest{
		Id: uint64(bookInput.Id),
	})
	if err != nil {
		return errors.Wrap(ErrValidation, "book with this id doesn't exist")
	}
	if bookInput.Title == "" {
		return errors.Wrap(ErrValidation, "field: [title] cannot be empty")
	}
	_, err = s.Repository.BookUpdate(ctx, &pb.BookUpdateRequest{
		Id:       uint64(bookInput.Id),
		Title:    bookInput.Title,
		AuthorId: uint64(bookInput.AuthorId),
	})
	return err
}

func (s *Implementation) DeleteBook(ctx context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	_, err := s.Repository.BookDelete(ctx, &pb.BookDeleteRequest{
		Id: uint64(id),
	})
	return err
}

func (s *Implementation) ListBooks(ctx context.Context, params models.ListInput) ([]models.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	if strings.ToLower(params.Order) == "desc" {
		params.Order = "DESC"
	} else {
		params.Order = "ASC"
	}
	bookListResponse, err := s.Repository.BookList(ctx, &pb.BookListRequest{
		Offset: params.Offset,
		Limit:  params.Limit,
		Order:  params.Order,
	})
	if err != nil {
		return nil, err
	}
	booksData := bookListResponse.GetBooks()
	books := make([]models.Book, len(booksData))
	for i, book := range booksData {
		books[i] = models.Book{
			Id:       uint(book.Id),
			Title:    book.Title,
			AuthorId: uint(book.AuthorId),
		}
	}
	return books, err
}
