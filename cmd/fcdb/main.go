package main

import (
	"context"
	"flag"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/w84thesun/film-camera-database/pkg/db"
	"github.com/w84thesun/film-camera-database/pkg/routes"
)

var (
	mongoDBURI = flag.String("mongo-uri", "mongodb://localhost:27017", "MongoDB URI")
	listenAddr = flag.String("listen-addr", ":8080", "Listen address")
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

	handler := routes.NewHandler(coll)

	router := mux.NewRouter()

	router.HandleFunc("/cameras", handler.GetCameras).Methods("GET")

	err = http.ListenAndServe(*listenAddr, router)
	if err != nil {
		panic(err)
	}
}
