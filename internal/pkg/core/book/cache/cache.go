package cache

import "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/core/book/models"

type Interface interface {
	Get(id uint) (models.Book, error)
	Add(book models.Book) error
	Update(book models.Book) error
	Delete(id uint) error
	List() []models.Book
}
