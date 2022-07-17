package handlers

import (
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/commander"

	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/storage"
)

const (
	helpCmd   = "help"
	listCmd   = "list"
	addCmd    = "add"
	removeCmd = "remove"
	editCmd   = "edit"
	readCmd   = "read"
)

var BadArgument = errors.New("bad argument")

func listFunc(s string) string {
	data := storage.List()
	if len(data) == 0 {
		return "book list is empty"
	}
	res := make([]string, 0, len(data))
	for _, v := range data {
		res = append(res, v.String())
	}
	return strings.Join(res, "\n")
}

func helpFunc(s string) string {
	return "/help - list commands\n" +
		"/list - list data\n" +
		"/add <title> <author> - add new book with title and author\n" +
		"/read <id> - mark book with this id as readed\n" +
		"/remove <id> - remove book with this id"
}

func addFunc(data string) string {
	log.Printf("add command param: %s", data)
	params := strings.Split(data, "|")
	if len(params) != 2 {
		return errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params).Error()
	}
	book, err := storage.NewBook(params[0], params[1], true)
	if err != nil {
		return err.Error()
	}
	err = storage.Add(book)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("book %v added", book)
}

func removeFunc(data string) string {
	log.Printf("delete command param: %s", data)
	id, err := parseId(data)
	if err != nil {
		return err.Error()
	}
	err = storage.Remove(id)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("book with id %d was removed", id)
}

func readFunc(data string) string {
	log.Printf("read command param: %s", data)
	id, err := parseId(data)
	if err != nil {
		return err.Error()
	}
	err = storage.MarkRead(id)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("book with id %d was mark as read", id)
}

func AddHandlers(c *commander.Commander) {
	c.RegisterHandler(helpCmd, helpFunc)
	c.RegisterHandler(listCmd, listFunc)
	c.RegisterHandler(addCmd, addFunc)
	c.RegisterHandler(removeCmd, removeFunc)
	c.RegisterHandler(readCmd, readFunc)
}
