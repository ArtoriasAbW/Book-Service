package update

import (
	"strconv"
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
	return "update"
}

func (c *command) Description() string {
	return "update book information"
}

func (c *command) Process(args string) string {
	params := strings.Split(args, "|")
	if len(params) != 4 {
		return "invalid args"
	}
	id, err := strconv.ParseUint(strings.Trim(params[0], " "), 10, 64)
	if err != nil {
		return "invalid id"
	}
	var unread bool
	switch strings.Trim(params[3], " ") {
	case "read":
		unread = false
	case "unread":
		unread = true
	default:
		return "invalid book status"
	}
	if err := c.book.Update(models.Book{
		Id:     uint(id),
		Title:  strings.Trim(params[1], " "),
		Author: strings.Trim(params[2], " "),
		Unread: unread,
	}); err != nil {
		if errors.Is(err, bookPkg.ErrValidation) {
			return "invalid args"
		}
		return "internal error"
	}
	return "success"
}
