package model

type Book struct {
	BookId int64  `json:"bookId" bson:"bookId"`
	Title  string `json:"title" bson:"title"`
	Author string `json:"author" bson:"author"`
	Year   int    `json:"year" bson:"year"`
}
