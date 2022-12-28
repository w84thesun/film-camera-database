package main

import (
	"context"
	"film-camera-database/pkg/db"
	"flag"
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

	coll, err := db.Collection(ctx, *mongoDBURI)
	if err != nil {
		panic(err)
	}

}
