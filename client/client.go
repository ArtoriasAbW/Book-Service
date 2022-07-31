package main

import (
	"context"
	"log"

	pb "gitlab.ozon.dev/ArtoriasAbW/homework-01/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/// Book test client
type bookClient struct {
	client pb.BookClient
}

func New(client pb.BookClient) bookClient {
	return bookClient{
		client: client,
	}
}

func (bc *bookClient) createBooks(ctx context.Context) error {
	_, err := bc.client.BookCreate(ctx, &pb.BookCreateRequest{
		Title:  "ABC",
		Author: "DEF",
		Unread: false,
	})
	if err != nil {
		return err
	}
	_, err = bc.client.BookCreate(ctx, &pb.BookCreateRequest{
		Title:  "ZXC",
		Author: "QWE",
		Unread: false,
	})
	return err
}

func (bc *bookClient) printBooks(ctx context.Context) {
	response, err := bc.client.BookList(ctx, &pb.BookListRequest{})
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("response: [%v]\n", response)
}

func (bc *bookClient) deleteBook(ctx context.Context, id uint64) error {
	_, err := bc.client.BookDelete(ctx, &pb.BookDeleteRequest{
		Id: id,
	})
	return err
}

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err.Error())
	}
	bc := New(pb.NewBookClient(conn))
	ctx := context.Background()
	if err = bc.createBooks(ctx); err != nil {
		log.Fatal(err.Error())
	}
	bc.printBooks(ctx)
	if err = bc.deleteBook(ctx, 1); err != nil {
		log.Fatal(err.Error())
	}
	bc.printBooks(ctx)
}
