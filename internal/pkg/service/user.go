package service

import (
	"context"
	"fmt"
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

func (s *service) AddUser(ctx context.Context, userInput models.User) error {
	fmt.Println("service:", userInput)
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	var user repoModels.User
	if userInput.Username == "" {
		return errors.Wrap(ErrValidation, "field: [title] cannot be empty")
	}
	user.Username = userInput.Username
	err := s.Repository.AddUser(ctx, user)
	return err
}