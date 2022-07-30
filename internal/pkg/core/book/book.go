package book

import (
	"github.com/pkg/errors"

	cachePkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/core/book/cache"
	localCachePkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/core/book/cache/local"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/core/book/models"
)

var ErrValidation = errors.New("invalid data")

type Interface interface {
	Create(book models.BookCreateInput) error
	Update(book models.Book) error
	Delete(id uint) error
	List() []models.Book
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

func (c *core) Create(bookInput models.BookCreateInput) error {
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
	err := c.cache.Add(book)
	if err != nil {
		return errors.Wrap(ErrValidation, err.Error())
	}
	return nil
}

func (c *core) Update(book models.Book) error {
	if book.Author == "" {
		return errors.Wrap(ErrValidation, "field: [author] cannot be empty")

	}
	if book.Title == "" {
		return errors.Wrap(ErrValidation, "field: [title] cannot be empty")
	}
	err := c.cache.Update(book)
	if err != nil {
		return errors.Wrap(ErrValidation, err.Error())
	}
	return nil
}

func (c *core) Delete(id uint) error {
	err := c.cache.Delete(id)
	if err != nil {
		return errors.Wrap(ErrValidation, err.Error())
	}
	return nil
}

func (c *core) List() []models.Book {
	return c.cache.List()
}
