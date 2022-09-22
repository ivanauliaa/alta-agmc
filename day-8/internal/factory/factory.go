package factory

import (
	"day6-hexagonal/database"
	"day6-hexagonal/internal/repository"
	"day6-hexagonal/mongo"
)

type Factory struct {
	UserRepository repository.UserInterface
	BookRepository repository.BookInterface
}

func NewFactory() *Factory {
	db := database.GetConnection()
	mongo := mongo.GetConnection()

	return &Factory{
		UserRepository: repository.NewUser(db),
		BookRepository: repository.NewBook(mongo),
	}
}
