package service

import (
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"
	repoModels "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/models"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/service/models"
)

func (s *service) GetAuthor(ctx context.Context, id uint) (models.Author, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*500))
	defer cancel()
	var err error
	author, err := s.Repository.GetAuthorById(ctx, id)
	return models.Author{
		Id:   author.Id,
		Name: author.Name,
	}, err
}

func (s *service) AddAuthor(ctx context.Context, authorInput models.Author) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	var author repoModels.Author
	if authorInput.Name == "" {
		return errors.Wrap(ErrValidation, "field: [name] cannot be empty")
	}
	author.Name = authorInput.Name
	err := s.Repository.AddAuthor(ctx, author)
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

func (s *service) ListAuthors(ctx context.Context, params models.ListInput) ([]models.Author, error) {
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
	authorsRepo, err := s.Repository.ListAuthors(ctx, repoParams)
	if err != nil {
		return nil, err
	}
	authors := make([]models.Author, len(authorsRepo))
	for i, author := range authorsRepo {
		authors[i] = models.Author{
			Id:   author.Id,
			Name: author.Name,
		}
	}
	return authors, nil
}
