package gomodel

import (
	"context"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var instance *mongo.Database
var ctx context.Context
var cancel context.CancelFunc
var mongoOnce sync.Once

func MongoInstance() (*mongo.Database, context.Context, *mongo.Client, context.CancelFunc) {
	mongoOnce.Do(func() {
		clientOption := options.Client().ApplyURI(os.Getenv("DATABASE_URI"))
		client, err := mongo.Connect(context.TODO(), clientOption)
		if err != nil {
			log.Fatal(err)
		}

		err = client.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal(err)
		}

		instance = client.Database(os.Getenv("DATABASE_NAME"))
	})

	return instance, ctx, client, cancel
}
