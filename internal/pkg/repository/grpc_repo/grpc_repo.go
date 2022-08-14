package grpc_repo

import (
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/pkg/api"
)

type repository struct {
	api.BookClient
	api.AuthorClient
	api.UserClient
	api.ReviewClient
}

func NewRepository(bookClient api.BookClient, authorClient api.AuthorClient, userClient api.UserClient, reviewClient api.ReviewClient) *repository {
	return &repository{
		BookClient:   bookClient,
		AuthorClient: authorClient,
		UserClient:   userClient,
		ReviewClient: reviewClient,
	}
}
