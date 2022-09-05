package main

import (
	"context"
	"net"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
	apiPkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/api"
	loggerPkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/logger"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/postgres"
	pb "gitlab.ozon.dev/ArtoriasAbW/homework-01/pkg/api"
	"google.golang.org/grpc"
)

func runGRPC(ctx context.Context) {
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		loggerPkg.Logger.Fatal(err.Error())
	}

	db := sqlx.MustOpen("postgres", os.Getenv("POSTGRES_CONNECTION"))
	//nolint:errcheck
	defer db.Close()
	if err != nil {
		loggerPkg.Logger.Fatal("can't connect to database: " + err.Error())
	}
	if err = db.PingContext(ctx); err != nil {
		loggerPkg.Logger.Fatal("ping database error: " + err.Error())
	}
	repo := postgres.NewRepository(db)

	grpcServer := grpc.NewServer()
	api := apiPkg.New(repo)
	pb.RegisterBookServer(grpcServer, api)
	pb.RegisterAuthorServer(grpcServer, api)
	pb.RegisterUserServer(grpcServer, api)
	pb.RegisterReviewServer(grpcServer, api)
	if err = grpcServer.Serve(listener); err != nil {
		loggerPkg.Logger.Fatal(err.Error())
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := loggerPkg.InitLogger()
	if err != nil {
		loggerPkg.Logger.Fatal("Failed to init logger")
	}
	if err := godotenv.Load(); err != nil {
		loggerPkg.Logger.Info("cannot load .env file")
	}
	defer loggerPkg.Logger.Sync()
	runGRPC(ctx)
}
