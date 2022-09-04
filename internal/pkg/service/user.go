package service

import (
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/service/models"
	pb "gitlab.ozon.dev/ArtoriasAbW/homework-01/pkg/api"
)

func (s *Implementation) GetUser(ctx context.Context, id uint) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*500))
	defer cancel()
	var err error
	userResponse, err := s.Repository.UserGet(ctx, &pb.UserGetRequest{
		Id: uint64(id),
	})
	if err != nil {
		return models.User{}, err
	}
	return models.User{
		Id:       uint(userResponse.Id),
		Username: userResponse.GetUsername(),
	}, nil
}

func (s *Implementation) AddUser(ctx context.Context, userInput models.User) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	if userInput.Username == "" {
		return 0, errors.Wrap(ErrValidation, "field: [username] cannot be empty")
	}
	userResponse, err := s.Repository.UserCreate(ctx, &pb.UserCreateRequest{
		Username: userInput.Username,
	})
	if err != nil {
		return 0, err
	}
	return userResponse.GetId(), nil
}

func (s *Implementation) DeleteUser(ctx context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	_, err := s.Repository.UserDelete(ctx, &pb.UserDeleteRequest{
		Id: uint64(id),
	})
	return err
}

func (s *Implementation) UpdateUser(ctx context.Context, userInput models.User) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	if userInput.Username == "" {
		return errors.Wrap(ErrValidation, "field: [username] cannot be empty")
	}
	_, err := s.Repository.UserGet(ctx, &pb.UserGetRequest{
		Id: uint64(userInput.Id),
	})
	if err != nil {
		return errors.Wrap(ErrValidation, "user with this id doesn't exist")
	}
	_, err = s.Repository.UserUpdate(ctx, &pb.UserUpdateRequest{
		Id:       uint64(userInput.Id),
		Username: userInput.Username,
	})
	return err
}

func (s *Implementation) ListUsers(ctx context.Context, params models.ListInput) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*1000))
	defer cancel()
	if strings.ToLower(params.Order) == "desc" {
		params.Order = "DESC"
	} else {
		params.Order = "ASC"
	}
	usersResponce, err := s.Repository.UserList(ctx, &pb.UserListRequest{
		Offset: params.Offset,
		Limit:  params.Limit,
		Order:  params.Order,
	})
	if err != nil {
		return nil, err
	}
	usersData := usersResponce.GetUsers()
	users := make([]models.User, len(usersData))
	for i, user := range usersData {
		users[i] = models.User{
			Id:       uint(user.Id),
			Username: user.Username,
		}
	}
	return users, nil
}
