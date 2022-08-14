package api

import (
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository"
	pb "gitlab.ozon.dev/ArtoriasAbW/homework-01/pkg/api"
)

func New(repo repository.Interface) Interface {
	return &handler{
		repo: repo,
	}
}

type Interface interface {
	pb.BookServer
	pb.AuthorServer
	pb.UserServer
	pb.ReviewServer
}

type handler struct {
	repo repository.Interface
	unimplementedServer
}

type unimplementedServer struct {
	pb.UnimplementedAuthorServer
	pb.UnimplementedBookServer
	pb.UnimplementedUserServer
	pb.UnimplementedReviewServer
}
