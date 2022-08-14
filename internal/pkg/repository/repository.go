package repository

import (
	pb "gitlab.ozon.dev/ArtoriasAbW/homework-01/pkg/api"
)

type Interface interface {
	book
	author
	user
	review
}

type book interface {
	pb.BookClient
}

type user interface {
	pb.UserClient
}

type author interface {
	pb.AuthorClient
}

type review interface {
	pb.ReviewClient
}
