package main

import (
	"context"
	"flag"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	mongoDBURI = flag.String("mongo-db-uri", "mongodb://localhost:27017", "MongoDB URI")
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if *mongoDBURI == "" {
		panic("mongo-db-uri is required")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(*mongoDBURI))
	if err != nil {
		panic(err)
	}

	client.Ping(ctx, nil)
}
