package service

import (
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"
	repoModels "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/models"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/service/models"
)

func (s *service) GetUser(ctx context.Context, id uint) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*500))
	defer cancel()
	var err error
	user, err := s.Repository.GetUserById(ctx, id)
	return models.User{
		Id:       user.Id,
		Username: user.Username,
	}, err
}

func (s *service) AddUser(ctx context.Context, userInput models.User) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	var user repoModels.User
	if userInput.Username == "" {
		return 0, errors.Wrap(ErrValidation, "field: [title] cannot be empty")
	}
	user.Username = userInput.Username
	id, err := s.Repository.AddUser(ctx, user)
	return id, err
}

func (s *service) DeleteUser(ctx context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	err := s.Repository.DeleteUser(ctx, id)
	return err
}

func (s *service) ListUsers(ctx context.Context, params models.ListInput) ([]models.User, error) {
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
	usersRepo, err := s.Repository.ListUsers(ctx, repoParams)
	if err != nil {
		return nil, err
	}
	users := make([]models.User, len(usersRepo))
	for i, user := range usersRepo {
		users[i] = models.User{
			Id:       user.Id,
			Username: user.Username,
		}
	}
	return users, nil
}
