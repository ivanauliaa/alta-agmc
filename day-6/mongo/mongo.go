package mongo

import (
	"context"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	db   *mongo.Database
	once sync.Once
	ctx  context.Context
)

func CreateConnection() {
	conf := dbConfig{
		User: os.Getenv("MONGO_USER"),
		Pass: os.Getenv("MONGO_PASS"),
		Host: os.Getenv("MONGO_HOST"),
		Port: os.Getenv("MONGO_PORT"),
		Name: os.Getenv("MONGO_NAME"),
	}

	mongoDb := mongoConfig{
		dbConfig: conf,
	}

	once.Do(func() {
		mongoDb.Connect()
	})
}

func GetConnection() *mongo.Database {
	if db == nil {
		CreateConnection()
	}

	return db
}
