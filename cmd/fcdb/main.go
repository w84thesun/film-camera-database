package main

import (
	"context"
	"flag"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	mongoDBURI = flag.String("mongo-uri", "mongodb://localhost:27017", "MongoDB URI")
)

func main() {
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if *mongoDBURI == "" {
		panic("mongo-db-uri is required")
	}

	println(*mongoDBURI)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(*mongoDBURI))
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	db := client.Database("film-cameras")

	coll := db.Collection("cameras")

	var doc bson.D
	res := coll.FindOne(ctx, bson.D{})
	if res.Err() != nil {
		panic(res.Err())
	}

	err = res.Decode(&doc)
	if err != nil {
		panic(err)
	}

	for key, value := range doc.Map() {
		println(key, value)
	}
}
