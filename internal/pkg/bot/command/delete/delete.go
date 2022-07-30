package delete

import (
	"log"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	commandPkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/bot/command"
	bookPkg "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/core/book"
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
	return "delete"
}

func (c *command) Description() string {
	return "delete book"
}

func (c *command) Process(args string) string {
	params := strings.Split(args, "|")
	if len(params) != 1 {
		return "invalid args"
	}
	id, err := strconv.ParseUint(strings.Trim(params[0], " "), 10, 64)
	if err != nil {
		return "invalid id"
	}
	if err := c.book.Delete(uint(id)); err != nil {
		log.Println(err.Error())
		if errors.Is(err, bookPkg.ErrValidation) {
			return "invalid args"
		}
		return "internal error"
	}
	return "success"
}
