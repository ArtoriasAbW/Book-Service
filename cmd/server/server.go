package main

import (
	"log"
	"net"

	apiPkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/api"
	bookPkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/core/book"
	pb "gitlab.ozon.dev/ArtoriasAbW/homework-01/pkg/api"
	"google.golang.org/grpc"
)

func runGrpc() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err.Error())
	}
	var book bookPkg.Interface
	{
		book = bookPkg.New()
	}
	grpcServer := grpc.NewServer()
	pb.RegisterBookServer(grpcServer, apiPkg.New(book))
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	runGrpc()
}
