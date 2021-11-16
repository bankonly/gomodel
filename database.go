package gomodel

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

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
		client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DATABASE_URI")))
		if err != nil {
			log.Fatal(err)
		}

		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}

		instance = client.Database(os.Getenv("DATABASE_NAME"))
	})

	return instance, ctx, client, cancel
}
