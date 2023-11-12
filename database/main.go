package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() *mongo.Client {
	options.ServerAPI(options.ServerAPIVersion1)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		panic(err)
	}

	client.Database()

	return client
}
