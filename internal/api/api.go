package services

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

func New(bookService service.Interface) pb.BookServer {
	return &implementation{
		bookService: bookService,
	}
}

type implementation struct {
	pb.UnimplementedBookServer
	bookService service.Interface
}

func (i *implementation) BookGet(_ context.Context, in *pb.BookGetRequest) (*pb.BookGetResponse, error) {
	book, err := i.bookService.GetBook(uint(in.Id))
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

func (i *implementation) BookCreate(_ context.Context, in *pb.BookCreateRequest) (*pb.BookCreateResponse, error) {
	fmt.Println("handler", in.GetTitle(), in.GetAuthorId())
	if err := i.bookService.AddBook(models.Book{
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

func (i *implementation) BookList(_ context.Context, in *pb.BookListRequest) (*pb.BookListResponse, error) {
	return &pb.BookListResponse{
		Books: nil,
	}, errors.New("not implemented")
}

func (i *implementation) BookUpdate(_ context.Context, in *pb.BookUpdateRequest) (*pb.BookUpdateResponse, error) {

	return &pb.BookUpdateResponse{}, nil
}
func (i *implementation) BookDelete(_ context.Context, in *pb.BookDeleteRequest) (*pb.BookDeleteResponse, error) {
	return &pb.BookDeleteResponse{}, nil
}
