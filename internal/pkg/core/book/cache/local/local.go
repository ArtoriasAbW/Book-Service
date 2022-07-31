package local

import (
	"sync"

	"github.com/pkg/errors"
	cachePkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/core/book/cache"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/core/book/models"
)

const poolSize = 10

var (
	ErrBookNotExists = errors.New("book doesn't exist")
	ErrBookExists    = errors.New("book already exists")
)

type cache struct {
	mu     sync.RWMutex
	data   map[uint]models.Book
	poolCh chan struct{}
}

func New() cachePkg.Interface {
	return &cache{
		mu:     sync.RWMutex{},
		data:   map[uint]models.Book{},
		poolCh: make(chan struct{}, poolSize),
	}
}

func (c *cache) Add(book models.Book) error {
	c.poolCh <- struct{}{}
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.data[book.Id]; ok {
		return errors.Wrapf(ErrBookExists, "book-id: [%d]", book.Id)
	}
	c.data[book.Id] = book
	<-c.poolCh
	return nil
}

func (c *cache) Update(book models.Book) error {
	c.poolCh <- struct{}{}
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.data[book.Id]; !ok {
		return errors.Wrapf(ErrBookNotExists, "book-id: [%d]", book.Id)
	}
	c.data[book.Id] = book
	<-c.poolCh
	return nil
}

func (c *cache) Delete(id uint) error {
	c.poolCh <- struct{}{}
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.data[id]; ok {
		delete(c.data, id)
		return nil
	}
	<-c.poolCh
	return errors.Wrapf(ErrBookNotExists, "book-id: [%d]", id)
}

func (c *cache) List() []models.Book {
	c.poolCh <- struct{}{}
	c.mu.RLock()
	defer c.mu.RUnlock()
	bookList := make([]models.Book, 0, len(c.data))
	for _, value := range c.data {
		bookList = append(bookList, value)
	}
	<-c.poolCh
	return bookList
}
