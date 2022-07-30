package storage

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

var lastId = uint(0)

type Book struct {
	id     uint
	title  string
	author string
	unread bool
}

func NewBook(title, author, status string) (*Book, error) {
	lastId++
	return CreateBook(lastId, title, author, status)
}

func CreateBook(id uint, title, author, status string) (*Book, error) {
	book := Book{}
	if err := book.SetAuthor(author); err != nil {
		return nil, err
	}
	if err := book.SetTitle(title); err != nil {
		return nil, err
	}
	if err := book.SetStatus(status); err != nil {
		return nil, err
	}
	book.id = id
	return &book, nil
}

func (b *Book) SetTitle(title string) error {
	b.title = strings.Trim(title, " ")
	return nil
}

func (b *Book) SetAuthor(author string) error {
	b.author = strings.Trim(author, " ")
	return nil
}

func (b *Book) SetStatus(status string) error {
	status = strings.Trim(status, " ")
	switch status {
	case "unread":
		b.unread = true
		return nil
	case "read":
		b.unread = false
		return nil
	}
	return errors.New("invalid book status: " + status)
}

func (b Book) String() string {
	var state string
	if b.unread {
		state = "unread"
	} else {
		state = "read"
	}
	return fmt.Sprintf("%d: %s | %s | %s", b.id, b.title, b.author, state)
}

func (b Book) GetTitle() string {
	return b.title
}

func (b Book) GetAuthor() string {
	return b.title
}

func (b Book) GetId() uint {
	return b.id
}
