package service

import (
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/service/models"
)

func New(d Deps) *service {
	return &service{Deps: d}
}

type Interface interface {
	GetBook(id uint) (models.Book, error)
	AddBook(bookInput models.Book) error
}

type service struct {
	Deps
}

type Deps struct {
	Repository repository.Interface
}
