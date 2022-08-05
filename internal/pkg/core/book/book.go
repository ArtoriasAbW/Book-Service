package book

import (
	"context"
	"time"

	"github.com/pkg/errors"

	cachePkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/core/book/cache"
	localCachePkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/core/book/cache/local"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/core/book/models"
)

var ErrValidation = errors.New("invalid data")

type Interface interface {
	Get(id uint) (models.Book, error)
	Create(book models.BookCreateInput) error
	Update(book models.Book) error
	Delete(id uint) error
	List() ([]models.Book, error)
}

func New() Interface {
	return &core{
		id:    0,
		cache: localCachePkg.New(),
	}
}

type core struct {
	id    uint
	cache cachePkg.Interface
}

func (c *core) Get(id uint) (models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*50))
	defer cancel()
	var err error
	var book models.Book
	operationDone := make(chan struct{})
	go func() {
		book, err = c.cache.Get(id)
		operationDone <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		return models.Book{}, ctx.Err()
	case <-operationDone:
		if err != nil {
			return models.Book{}, errors.Wrap(ErrValidation, err.Error())
		}
		return book, nil
	}
}

func (c *core) Create(bookInput models.BookCreateInput) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*50))
	defer cancel()
	var book models.Book
	if bookInput.Author == "" {
		return errors.Wrap(ErrValidation, "field: [author] cannot be empty")
	}
	book.Author = bookInput.Author
	if bookInput.Title == "" {
		return errors.Wrap(ErrValidation, "field: [title] cannot be empty")
	}
	book.Title = bookInput.Title
	book.Unread = bookInput.Unread
	c.id++
	book.Id = c.id
	var err error
	operationDone := make(chan struct{})
	go func() {
		err = c.cache.Add(book)
		operationDone <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-operationDone:
		if err != nil {
			return errors.Wrap(ErrValidation, err.Error())
		}
		return nil
	}
}

func (c *core) Update(book models.Book) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*50))
	defer cancel()
	if book.Author == "" {
		return errors.Wrap(ErrValidation, "field: [author] cannot be empty")

	}
	if book.Title == "" {
		return errors.Wrap(ErrValidation, "field: [title] cannot be empty")
	}
	var err error
	operationDone := make(chan struct{})
	go func() {
		err = c.cache.Update(book)
		operationDone <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-operationDone:
		if err != nil {
			return errors.Wrap(ErrValidation, err.Error())
		}
		return nil
	}
}

func (c *core) Delete(id uint) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*50))
	defer cancel()
	var err error
	operationDone := make(chan struct{})
	go func() {
		err = c.cache.Delete(id)
		operationDone <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-operationDone:
		if err != nil {
			return errors.Wrap(ErrValidation, err.Error())
		}
		return nil
	}
}

func (c *core) List() ([]models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*50))
	defer cancel()
	var books []models.Book
	operationDone := make(chan struct{})
	go func() {
		books = c.cache.List()
		operationDone <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		return []models.Book{}, ctx.Err()
	case <-operationDone:
		return books, nil
	}
}
