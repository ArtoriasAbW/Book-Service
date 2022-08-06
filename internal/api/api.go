package api

import (
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/service"
	pb "gitlab.ozon.dev/ArtoriasAbW/homework-01/pkg/api"
)

func New(service service.Interface) Interface {
	return &handler{
		service: service,
	}
}

type Interface interface {
	pb.BookServer
	pb.AuthorServer
	pb.UserServer
}

type handler struct {
	service service.Interface
	unimplementedServer
}

type unimplementedServer struct {
	pb.UnimplementedAuthorServer
	pb.UnimplementedBookServer
	pb.UnimplementedUserServer
}
