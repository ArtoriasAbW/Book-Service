package service

import (
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/service/models"
	pb "gitlab.ozon.dev/ArtoriasAbW/homework-01/pkg/api"
)

func (s *Implementation) GetAuthor(ctx context.Context, id uint) (models.Author, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*500))
	defer cancel()
	var err error
	authorGetResponce, err := s.Repository.AuthorGet(ctx, &pb.AuthorGetRequest{
		Id: uint64(id),
	})
	if err != nil {
		return models.Author{}, err
	}
	return models.Author{
		Id:   uint(authorGetResponce.GetId()),
		Name: authorGetResponce.GetName(),
	}, nil
}

func (s *Implementation) AddAuthor(ctx context.Context, authorInput models.Author) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	if authorInput.Name == "" {
		return 0, errors.Wrap(ErrValidation, "field: [name] cannot be empty")
	}
	authorCreateResponce, err := s.Repository.AuthorCreate(ctx, &pb.AuthorCreateRequest{
		Name: authorInput.Name,
	})
	if err != nil {
		return 0, err
	}
	return authorCreateResponce.GetId(), nil
}

func (s *Implementation) DeleteAuthor(ctx context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	_, err := s.Repository.AuthorDelete(ctx, &pb.AuthorDeleteRequest{
		Id: uint64(id),
	})
	return err
}

func (s *Implementation) UpdateAuthor(ctx context.Context, authorInput models.Author) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	_, err := s.Repository.AuthorGet(ctx, &pb.AuthorGetRequest{
		Id: uint64(authorInput.Id),
	})
	if err != nil {
		return errors.Wrap(ErrValidation, "author with this id doesn't exist")
	}
	if authorInput.Name == "" {
		return errors.Wrap(ErrValidation, "field: [name] cannot be empty")
	}
	_, err = s.Repository.AuthorUpdate(ctx, &pb.AuthorUpdateRequest{
		Id:   uint64(authorInput.Id),
		Name: authorInput.Name,
	})
	return err
}

func (s *Implementation) ListAuthors(ctx context.Context, params models.ListInput) ([]models.Author, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	if strings.ToLower(params.Order) == "desc" {
		params.Order = "DESC"
	} else {
		params.Order = "ASC"
	}
	authorListResponce, err := s.Repository.AuthorList(ctx, &pb.AuthorListRequest{
		Limit:  params.Limit,
		Offset: params.Offset,
		Order:  params.Order,
	})
	if err != nil {
		return nil, err
	}
	authorsData := authorListResponce.GetAuthors()
	authors := make([]models.Author, len(authorsData))
	for i, author := range authorsData {
		authors[i] = models.Author{
			Id:   uint(author.Id),
			Name: author.Name,
		}
	}
	return authors, err

}
