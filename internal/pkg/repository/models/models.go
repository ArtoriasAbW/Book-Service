package models

import "time"

type User struct {
	Id       uint   `db:"id"`
	Username string `db:"username"`
}

type Author struct {
	Id   uint   `db:"id"`
	Name string `db:"author"`
}

type Book struct {
	Id       uint   `db:"id"`
	Title    string `db:"title"`
	AuthorId uint   `db:"author_id"`
}

type Review struct {
	Id         uint      `db:"id"`
	Rating     uint      `db:"rating"`
	ReviewText string    `db:"review_text"`
	Time       time.Time `db:"time"`
	BookId     uint      `db:"book_id"`
	UserId     uint      `db:"user_id"`
}
