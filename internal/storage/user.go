package storage

import (
	"fmt"
)

var lastId = uint(0)

type Book struct {
	id     uint
	title  string
	author string
	unread bool
}

func NewBook(title, author string, isReaded bool) (*Book, error) {
	book := Book{}
	if err := book.SetAuthor(author); err != nil {
		return nil, err
	}
	if err := book.SetTitle(title); err != nil {
		return nil, err
	}
	if err := book.SetStatus(isReaded); err != nil {
		return nil, err
	}
	lastId++
	book.id = lastId
	return &book, nil
}

func (b *Book) SetTitle(title string) error {
	b.title = title
	return nil
}

func (b *Book) SetAuthor(author string) error {
	b.author = author
	return nil
}

func (b *Book) SetStatus(isReaded bool) error {
	b.unread = isReaded
	return nil
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
