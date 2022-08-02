package api

import (
	"context"
	"errors"

	bookPkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/core/book"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/core/book/models"
	pb "gitlab.ozon.dev/ArtoriasAbW/homework-01/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func New(book bookPkg.Interface) pb.BookServer {
	return &implementation{
		book: book,
	}
}

type implementation struct {
	pb.UnimplementedBookServer
	book bookPkg.Interface
}

func (i *implementation) BookGet(_ context.Context, in *pb.BookGetRequest) (*pb.BookGetResponse, error) {
	book, err := i.book.Get(uint(in.Id))
	if err != nil {
		if errors.Is(err, bookPkg.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.BookGetResponse{
		Title:  book.Title,
		Author: book.Author,
		Unread: book.Unread,
	}, nil
}

func (i *implementation) BookCreate(_ context.Context, in *pb.BookCreateRequest) (*pb.BookCreateResponse, error) {
	if err := i.book.Create(models.BookCreateInput{
		Title:  in.GetTitle(),
		Author: in.GetAuthor(),
		Unread: in.GetUnread(),
	}); err != nil {
		if errors.Is(err, bookPkg.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.BookCreateResponse{}, nil

}

func (i *implementation) BookList(_ context.Context, in *pb.BookListRequest) (*pb.BookListResponse, error) {
	books, err := i.book.List()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	result := make([]*pb.BookListResponse_Book, 0, len(books))
	for _, book := range books {
		result = append(result, &pb.BookListResponse_Book{
			Id:     uint64(book.Id),
			Title:  book.Title,
			Author: book.Author,
			Unread: book.Unread,
		})
	}
	return &pb.BookListResponse{
		Books: result,
	}, nil
}

func (i *implementation) BookUpdate(_ context.Context, in *pb.BookUpdateRequest) (*pb.BookUpdateResponse, error) {
	if err := i.book.Update(
		models.Book{
			Id:     uint(in.Id),
			Title:  in.Title,
			Author: in.Author,
			Unread: in.Unread,
		}); err != nil {
		if errors.Is(err, bookPkg.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.BookUpdateResponse{}, nil
}
func (i *implementation) BookDelete(_ context.Context, in *pb.BookDeleteRequest) (*pb.BookDeleteResponse, error) {
	if err := i.book.Delete(uint(in.GetId())); err != nil {
		if errors.Is(err, bookPkg.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.BookDeleteResponse{}, nil
}
