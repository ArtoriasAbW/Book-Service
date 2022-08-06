package models

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
	Id         uint   `db:"id"`
	Rating     uint   `db:"rating"`
	ReviewText string `db:"review_text"`
	Time       uint   `db:"time"`
	BookId     uint   `db:"book_id"`
	UserId     uint   `db:"user_id"`
}

type ReviewListInput struct {
	Limit  uint32
	Offset uint32
	Order  string
}
