package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
	Year   int    `json:"year" validate:"required"`
}
