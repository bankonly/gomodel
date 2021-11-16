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
var mongoOnce sync.Once

func MongoInstance() (*mongo.Database, *mongo.Client, context.Context) {
	mongoOnce.Do(func() {
		ctx = context.TODO()
		clientOption := options.Client().ApplyURI(os.Getenv("DATABASE_URI"))
		client, err := mongo.Connect(ctx, clientOption)
		if err != nil {
			log.Fatal(err)
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}

		instance = client.Database(os.Getenv("DATABASE_NAME"))
	})

	return instance, client, ctx
}
