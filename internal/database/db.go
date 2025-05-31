package database

import (
	"context"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance    *mongo.Client
	clientInstanceErr error
	mongoOnce         sync.Once
	connectionTimeout = 10 * time.Second
)

func getClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), connectionTimeout)
		defer cancel()

		clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
		clientInstance, clientInstanceErr = mongo.Connect(ctx, clientOptions)
	})
	return clientInstance, clientInstanceErr
}

func GetMongoCollection(database, collection string) *mongo.Collection {
	client, err := getClient()
	if err != nil {
		panic(err)
	}
	return client.Database(database).Collection(collection)
}
