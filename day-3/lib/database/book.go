package database

import (
	"day3-middleware/config"
	"day3-middleware/models"
)

func GetBooks() (interface{}, error) {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func CreateBook(book models.Book) (interface{}, error) {
	if err := config.DB.Create(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func GetBookById(id string) (interface{}, error) {
	var book models.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func UpdateBookById(id string, bookUpdate models.Book) (interface{}, error) {
	var book models.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		return nil, err
	}

	config.DB.Model(&book).Updates(bookUpdate)

	return book, nil
}

func DeleteBookById(id string) error {
	if err := config.DB.Delete(&models.Book{}, id).Error; err != nil {
		return err
	}

	return nil
}
