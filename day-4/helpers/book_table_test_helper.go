package helpers

import (
	"day3-middleware/config"
	"day3-middleware/models"

	"gorm.io/gorm"
)

func AddBookWithId(id uint) {
	book := models.Book{
		Model: gorm.Model{
			ID: id,
		},
		Title:  "agmc",
		Author: "alta",
		Year:   2022,
	}

	config.DB.Create(&book)
}

func CleanBooksTable() {
	config.DB.Exec("DELETE FROM books")
}
