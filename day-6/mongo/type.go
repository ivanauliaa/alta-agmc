package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	dbConfig struct {
		User string
		Pass string
		Host string
		Port string
		Name string
	}

	mongoConfig struct {
		dbConfig
	}
)

func (conf *mongoConfig) Connect() {
	dsn := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		conf.User,
		conf.Pass,
		conf.Host,
		conf.Port,
	)

	var err error

	clientOptions := options.Client()
	clientOptions.ApplyURI(dsn)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		panic(err)
	}

	ctx = context.TODO()

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	db = client.Database(conf.Name)
}
