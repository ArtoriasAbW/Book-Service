package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"

	apiPkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/api"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/config"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/postgres"
	servicePkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/service"
	pb "gitlab.ozon.dev/ArtoriasAbW/homework-01/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func runGRPC() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err.Error())
	}

	psqlConn := fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=book_service sslmode=disable", config.GetPostgresUser(), config.GetPostgresPassword())
	pool, err := pgxpool.Connect(ctx, psqlConn)
	if err != nil {
		log.Fatal("can't connect to database", err)
	}
	defer pool.Close()
	if err := pool.Ping(ctx); err != nil {
		log.Fatal("ping database error", err)
	}
	repo := postgres.NewRepository(pool)

	service := servicePkg.New(servicePkg.Deps{Repository: repo})
	grpcServer := grpc.NewServer()
	api := apiPkg.New(service)
	pb.RegisterBookServer(grpcServer, api)
	pb.RegisterAuthorServer(grpcServer, api)
	pb.RegisterUserServer(grpcServer, api)
	pb.RegisterReviewServer(grpcServer, api)
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err.Error())
	}
}

func runREST() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(headerMatcherREST),
	)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterBookHandlerFromEndpoint(ctx, mux, ":8081", opts); err != nil {
		log.Fatal(err.Error())
	}
	if err := pb.RegisterAuthorHandlerFromEndpoint(ctx, mux, ":8081", opts); err != nil {
		log.Fatal(err.Error())
	}
	if err := pb.RegisterUserHandlerFromEndpoint(ctx, mux, ":8081", opts); err != nil {
		log.Fatal(err.Error())
	}
	if err := pb.RegisterReviewHandlerFromEndpoint(ctx, mux, ":8081", opts); err != nil {
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
	if err := godotenv.Load(); err != nil {
		log.Println("cannot load .env file")
	}
	go runREST()
	runGRPC()
}
