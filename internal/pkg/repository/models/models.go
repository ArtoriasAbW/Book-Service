package models

import "time"

type User struct {
	ID       uint   `db:"id"`
	Username string `db:"username"`
}

type Author struct {
	ID   uint   `db:"id"`
	Name string `db:"author"`
}

type Book struct {
	ID       uint   `db:"id"`
	Title    string `db:"title"`
	AuthorID uint   `db:"author_id"`
}

type Review struct {
	ID         uint      `db:"id"`
	Rating     uint      `db:"rating"`
	ReviewText string    `db:"review_text"`
	Time       time.Time `db:"time"`
	BookID     uint      `db:"book_id"`
	UserID     uint      `db:"user_id"`
}

type ListInput struct {
	Limit  uint32
	Offset uint32
	Order  string
}
