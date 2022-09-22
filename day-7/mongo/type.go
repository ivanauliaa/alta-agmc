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
	dsn := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority",
		conf.User,
		conf.Pass,
		conf.Host,
	)

	var err error

	serverApiOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client()
	clientOptions.ApplyURI(dsn).SetServerAPIOptions(serverApiOptions)
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
