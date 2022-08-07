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

func (h *handler) BookGet(ctx context.Context, in *pb.BookGetRequest) (*pb.BookGetResponse, error) {
	book, err := h.service.GetBook(ctx, uint(in.Id))
	if err != nil {
		if errors.Is(err, service.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.BookGetResponse{
		Title:    book.Title,
		AuthorId: uint64(book.AuthorId),
	}, nil
}

func (h *handler) BookCreate(ctx context.Context, in *pb.BookCreateRequest) (*pb.BookCreateResponse, error) {
	if err := h.service.AddBook(ctx, models.Book{
		Title:    in.GetTitle(),
		AuthorId: uint(in.GetAuthorId()),
	}); err != nil {
		if errors.Is(err, service.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.BookCreateResponse{}, nil

}

func (h *handler) BookList(ctx context.Context, in *pb.BookListRequest) (*pb.BookListResponse, error) {
	return &pb.BookListResponse{
		Books: nil,
	}, errors.New("not implemented")
}

func (h *handler) BookUpdate(ctx context.Context, in *pb.BookUpdateRequest) (*pb.BookUpdateResponse, error) {

	return &pb.BookUpdateResponse{}, nil
}
func (h *handler) BookDelete(ctx context.Context, in *pb.BookDeleteRequest) (*pb.BookDeleteResponse, error) {
	if err := h.service.DeleteBook(ctx, uint(in.GetId())); err != nil {
		if errors.Is(err, service.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.BookDeleteResponse{}, nil
}
