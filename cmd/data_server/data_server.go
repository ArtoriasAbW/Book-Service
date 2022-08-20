package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
	apiPkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/api"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/postgres"
	pb "gitlab.ozon.dev/ArtoriasAbW/homework-01/pkg/api"
	"google.golang.org/grpc"
)

func runGRPC(ctx context.Context) {
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal(err.Error())
	}

	db := sqlx.MustOpen("postgres", os.Getenv("POSTGRES_CONNECTION"))
	defer db.Close()
	if err != nil {
		log.Fatal("can't connect to database", err)
	}
	if err = db.PingContext(ctx); err != nil {
		log.Fatal("ping database error", err)
	}
	repo := postgres.NewRepository(db)

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
