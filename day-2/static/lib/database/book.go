package database

import (
	"day2-static/models"
	"fmt"
	"time"
)

var books []models.Book = []models.Book{}

func GetBooks() (interface{}, error) {
	return books, nil
}

func CreateBook(book models.Book) (interface{}, error) {
	book.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	books = append(books, book)

	return book, nil
}

func GetBookById(id string) (interface{}, error) {
	for index := range books {
		if books[index].ID == id {
			return books[index], nil
		}
	}

	return nil, fmt.Errorf("book not found")
}

func UpdateBookById(id string, bookUpdate models.Book) (interface{}, error) {
	for index := range books {
		if books[index].ID == id {
			books[index].Title = bookUpdate.Title
			books[index].Author = bookUpdate.Author
			books[index].Year = bookUpdate.Year

			return books[index], nil
		}
	}

	return nil, fmt.Errorf("book not found")
}

func DeleteBookById(id string) error {
	for index := range books {
		if books[index].ID == id {
			books = append(books[:index], books[index+1:]...)

			return nil
		}
	}

	return fmt.Errorf("book not found")
}
