package gomodel

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoInstance() (*mongo.Client, context.Context, context.CancelFunc) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DATABASE_URI")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client, ctx, cancel
}

func MongoWithDBName(database string) (*mongo.Client, context.Context, context.CancelFunc) {
	client, ctx, cancel := MongoInstance()
	client.Database(database)
	return client, ctx, cancel
}

func Mongo() (*mongo.Database, *mongo.Client, context.Context, context.CancelFunc) {
	client, ctx, cancel := MongoInstance()
	instance := client.Database(os.Getenv("DATABASE_NAME"))
	return instance, client, ctx, cancel
}
