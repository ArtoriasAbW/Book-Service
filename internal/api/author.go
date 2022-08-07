package api

import (
	"context"
	"errors"

	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/service"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/service/models"
	pb "gitlab.ozon.dev/ArtoriasAbW/homework-01/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) AuthorGet(ctx context.Context, in *pb.AuthorGetRequest) (*pb.AuthorGetResponse, error) {
	author, err := h.service.GetAuthor(ctx, uint(in.GetId()))
	if err != nil {
		if errors.Is(err, service.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.AuthorGetResponse{
		Id:   uint64(author.Id),
		Name: author.Name,
	}, nil
}

func (h *handler) AuthorCreate(ctx context.Context, in *pb.AuthorCreateRequest) (*pb.AuthorCreateResponse, error) {
	if err := h.service.AddAuthor(ctx, models.Author{
		Name: in.GetName(),
	}); err != nil {
		if errors.Is(err, service.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.AuthorCreateResponse{}, nil

}

func (h *handler) AuthorList(ctx context.Context, in *pb.AuthorListRequest) (*pb.AuthorListResponse, error) {
	return &pb.AuthorListResponse{
		Authors: nil,
	}, errors.New("not implemented")
}

func (h *handler) AuthorUpdate(_ context.Context, in *pb.AuthorUpdateRequest) (*pb.AuthorUpdateResponse, error) {

	return &pb.AuthorUpdateResponse{}, nil
}
func (h *handler) AuthorDelete(ctx context.Context, in *pb.AuthorDeleteRequest) (*pb.AuthorDeleteResponse, error) {
	if err := h.service.DeleteAuthor(ctx, uint(in.GetId())); err != nil {
		if errors.Is(err, service.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.AuthorDeleteResponse{}, nil
}
