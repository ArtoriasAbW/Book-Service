package api

import (
	"context"

	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/models"
	pb "gitlab.ozon.dev/ArtoriasAbW/homework-01/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) BookGet(ctx context.Context, in *pb.BookGetRequest) (*pb.BookGetResponse, error) {
	book, err := h.repo.GetBookById(ctx, uint(in.Id))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.BookGetResponse{
		Title:    book.Title,
		AuthorId: uint64(book.AuthorId),
	}, nil
}

func (h *handler) BookCreate(ctx context.Context, in *pb.BookCreateRequest) (*pb.BookCreateResponse, error) {
	id, err := h.repo.AddBook(ctx, models.Book{
		Title:    in.GetTitle(),
		AuthorId: uint(in.GetAuthorId()),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.BookCreateResponse{Id: id}, nil

}

func (h *handler) BookList(ctx context.Context, in *pb.BookListRequest) (*pb.BookListResponse, error) {
	books, err := h.repo.ListBooks(ctx, models.ListInput{
		Limit:  in.GetLimit(),
		Offset: in.GetOffset(),
		Order:  in.GetOrder(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	result := make([]*pb.BookListResponse_Book, 0, len(books))
	for _, book := range books {
		result = append(result, &pb.BookListResponse_Book{
			Id:       uint64(book.Id),
			Title:    book.Title,
			AuthorId: uint64(book.AuthorId),
		})
	}
	return &pb.BookListResponse{
		Books: result,
	}, nil
}

func (h *handler) BookUpdate(ctx context.Context, in *pb.BookUpdateRequest) (*pb.BookUpdateResponse, error) {
	if err := h.repo.UpdateBook(ctx, models.Book{
		Id:       uint(in.GetId()),
		Title:    in.GetTitle(),
		AuthorId: uint(in.GetAuthorId()),
	}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.BookUpdateResponse{}, nil
}
func (h *handler) BookDelete(ctx context.Context, in *pb.BookDeleteRequest) (*pb.BookDeleteResponse, error) {
	if err := h.repo.DeleteBook(ctx, uint(in.GetId())); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.BookDeleteResponse{}, nil
}
