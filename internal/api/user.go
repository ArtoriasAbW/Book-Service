package api

import (
	"context"
	"errors"
	"fmt"

	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/service"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/service/models"
	pb "gitlab.ozon.dev/ArtoriasAbW/homework-01/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) UserGet(ctx context.Context, in *pb.UserGetRequest) (*pb.UserGetResponse, error) {
	user, err := h.service.GetUser(ctx, uint(in.GetId()))
	if err != nil {
		if errors.Is(err, service.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.UserGetResponse{
		Id:       uint64(user.Id),
		Username: user.Username,
	}, nil
}
func (h *handler) UserCreate(ctx context.Context, in *pb.UserCreateRequest) (*pb.UserCreateResponse, error) {
	fmt.Println("create")
	id, err := h.service.AddUser(ctx, models.User{
		Username: in.GetUsername(),
	})
	if err != nil {
		if errors.Is(err, service.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.UserCreateResponse{Id: id}, nil
}
func (h *handler) UserList(ctx context.Context, in *pb.UserListRequest) (*pb.UserListResponse, error) {
	users, err := h.service.ListUsers(ctx, models.ListInput{
		Limit:  in.GetLimit(),
		Offset: in.GetOffset(),
		Order:  in.GetOrder(),
	})
	if err != nil {
		if errors.Is(err, service.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	result := make([]*pb.UserListResponse_User, 0, len(users))
	for _, user := range users {
		result = append(result, &pb.UserListResponse_User{
			Id:       uint64(user.Id),
			Username: user.Username,
		})
	}
	return &pb.UserListResponse{
		Users: result,
	}, nil
}
func (h *handler) UserUpdate(ctx context.Context, in *pb.UserUpdateRequest) (*pb.UserUpdateResponse, error) {
	if err := h.service.UpdateUser(ctx, models.User{
		Id:       uint(in.GetId()),
		Username: in.GetUsername(),
	}); err != nil {
		if errors.Is(err, service.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.UserUpdateResponse{}, nil
}
func (h *handler) UserDelete(ctx context.Context, in *pb.UserDeleteRequest) (*pb.UserDeleteResponse, error) {
	err := h.service.DeleteUser(ctx, uint(in.GetId()))
	if err != nil {
		if errors.Is(err, service.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.UserDeleteResponse{}, nil
}
