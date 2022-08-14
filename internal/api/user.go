package api

import (
	"context"

	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/models"
	pb "gitlab.ozon.dev/ArtoriasAbW/homework-01/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) UserGet(ctx context.Context, in *pb.UserGetRequest) (*pb.UserGetResponse, error) {
	user, err := h.repo.GetUserById(ctx, uint(in.GetId()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.UserGetResponse{
		Id:       uint64(user.Id),
		Username: user.Username,
	}, nil
}
func (h *handler) UserCreate(ctx context.Context, in *pb.UserCreateRequest) (*pb.UserCreateResponse, error) {
	id, err := h.repo.AddUser(ctx, models.User{
		Username: in.GetUsername(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.UserCreateResponse{Id: id}, nil
}
func (h *handler) UserList(ctx context.Context, in *pb.UserListRequest) (*pb.UserListResponse, error) {
	users, err := h.repo.ListUsers(ctx, models.ListInput{
		Limit:  in.GetLimit(),
		Offset: in.GetOffset(),
		Order:  in.GetOrder(),
	})
	if err != nil {
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
	if err := h.repo.UpdateUser(ctx, models.User{
		Id:       uint(in.GetId()),
		Username: in.GetUsername(),
	}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.UserUpdateResponse{}, nil
}
func (h *handler) UserDelete(ctx context.Context, in *pb.UserDeleteRequest) (*pb.UserDeleteResponse, error) {
	err := h.repo.DeleteUser(ctx, uint(in.GetId()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.UserDeleteResponse{}, nil
}
