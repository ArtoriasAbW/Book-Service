package list

import (
	"fmt"
	"strings"

	commandPkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/bot/command"
	bookPkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/core/book"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/core/book/models"
)

func New(book bookPkg.Interface) commandPkg.Interface {
	return &command{
		book: book,
	}
}

type command struct {
	book bookPkg.Interface
}

func (c *command) Name() string {
	return "list"
}

func (c *command) Description() string {
	return "list all books"
}

func (c *command) Process(_ string) string {
	bookList := c.book.List()
	if len(bookList) == 0 {
		return "no books"
	}
	result := make([]string, len(bookList))
	for i, book := range bookList {
		result[i] = bookToString(book)
	}
	return strings.Join(result, "\n")
}

func bookToString(book models.Book) string {
	var state string
	if book.Unread {
		state = "unread"
	} else {
		state = "read"
	}
	return fmt.Sprintf("%d: %s | %s | %s", book.Id, book.Title, book.Author, state)
}
