package service

import (
	"context"
	"time"

	"github.com/pkg/errors"
	repoModels "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/models"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/service/models"
)

func (c *service) GetAuthor(ctx context.Context, id uint) (models.Author, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*500))
	defer cancel()
	var err error
	author, err := c.Repository.GetAuthorById(ctx, id)
	return models.Author{
		Id:   author.Id,
		Name: author.Name,
	}, err
}

func (c *service) AddAuthor(ctx context.Context, authorInput models.Author) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	var author repoModels.Author
	if authorInput.Name == "" {
		return errors.Wrap(ErrValidation, "field: [name] cannot be empty")
	}
	author.Name = authorInput.Name
	err := c.Repository.AddAuthor(ctx, author)
	return err
}

func (c *service) DeleteAuthor(ctx context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	err := c.Repository.DeleteAuthor(ctx, id)
	return err
}

func (c *service) UpdateAuthor(id uint) error {
	return errors.New("not implemented")

}

func (c *service) ListAuthors() ([]models.Book, error) {
	return nil, errors.New("not implemented")
}
