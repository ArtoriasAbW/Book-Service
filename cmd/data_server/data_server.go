package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"

	apiPkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/api"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/config"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/postgres"
	pb "gitlab.ozon.dev/ArtoriasAbW/homework-01/pkg/api"
	"google.golang.org/grpc"
)

func runGRPC(ctx context.Context) {
	listener, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatal(err.Error())
	}

	psqlConn := fmt.Sprintf("host=localhost port=6432 user=%s password=%s dbname=book_service sslmode=disable", config.GetPostgresUser(), config.GetPostgresPassword())
	pool, err := pgxpool.Connect(ctx, psqlConn)
	if err != nil {
		log.Fatal("can't connect to database", err)
	}
	defer pool.Close()
	if err = pool.Ping(ctx); err != nil {
		log.Fatal("ping database error", err)
	}
	repo := postgres.NewRepository(pool)
	grpcServer := grpc.NewServer()
	api := apiPkg.New(repo)
	pb.RegisterBookServer(grpcServer, api)
	pb.RegisterAuthorServer(grpcServer, api)
	pb.RegisterUserServer(grpcServer, api)
	pb.RegisterReviewServer(grpcServer, api)
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := godotenv.Load(); err != nil {
		log.Println("cannot load .env file")
	}
	runGRPC(ctx)
}
