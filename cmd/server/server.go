package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	apiPkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/api"
	bookPkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/core/book"
	pb "gitlab.ozon.dev/ArtoriasAbW/homework-01/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func runGRPC() {
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

func runREST() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(headerMatcherREST),
	)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterBookHandlerFromEndpoint(ctx, mux, ":8081", opts); err != nil {
		log.Fatal(err.Error())
	}

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err.Error())
	}

}

func headerMatcherREST(key string) (string, bool) {
	switch key {
	case "Custom":
		return key, true
	default:
		return key, false
	}
}

func main() {
	
	go runREST()
	runGRPC()
}
