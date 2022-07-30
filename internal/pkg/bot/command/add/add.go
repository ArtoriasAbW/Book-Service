package add

import (
	"strings"

	"github.com/pkg/errors"
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
	return "add"
}

func (c *command) Description() string {
	return "add new book"
}

func (c *command) Process(args string) string {
	params := strings.Split(args, "|")
	if len(params) != 3 {
		return "invalid args"
	}
	var unread bool
	switch strings.Trim(params[2], " ") {
	case "read":
		unread = false
	case "unread":
		unread = true
	default:
		return "invalid book status"
	}
	if err := c.book.Create(models.BookCreateInput{
		Title:  strings.Trim(params[0], " "),
		Author: strings.Trim(params[1], " "),
		Unread: unread,
	}); err != nil {
		if errors.Is(err, bookPkg.ErrValidation) {
			return "invalid args"
		}
		return "internal error"
	}
	return "success"
}
