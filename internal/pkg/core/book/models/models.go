package models

type Book struct {
	Id     uint
	Title  string
	Author string
	Unread bool
}

type BookCreateInput struct {
	Title  string
	Author string
	Unread bool
}
