package gomodel

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoInstance() (*mongo.Database, *mongo.Client, context.Context, context.CancelFunc) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DATABASE_URI")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	instance := client.Database(os.Getenv("DATABASE_NAME"))

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return instance, client, ctx, cancel
}

// Update
