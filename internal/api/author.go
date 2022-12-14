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
	id, err := h.service.AddAuthor(ctx, models.Author{
		Name: in.GetName(),
	})
	if err != nil {
		if errors.Is(err, service.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.AuthorCreateResponse{Id: id}, nil

}

func (h *handler) AuthorList(ctx context.Context, in *pb.AuthorListRequest) (*pb.AuthorListResponse, error) {
	authors, err := h.service.ListAuthors(ctx, models.ListInput{
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

	result := make([]*pb.AuthorListResponse_Author, 0, len(authors))
	for _, author := range authors {
		result = append(result, &pb.AuthorListResponse_Author{
			Id:   uint64(author.Id),
			Name: author.Name,
		})
	}
	return &pb.AuthorListResponse{
		Authors: result,
	}, nil
}

func (h *handler) AuthorUpdate(ctx context.Context, in *pb.AuthorUpdateRequest) (*pb.AuthorUpdateResponse, error) {
	if err := h.service.UpdateAuthor(ctx, models.Author{
		Id:   uint(in.GetId()),
		Name: in.GetName(),
	}); err != nil {
		if errors.Is(err, service.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
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
