package routes

import "go.mongodb.org/mongo-driver/mongo"

type Handler struct {
	coll *mongo.Collection
}

func NewHandler(coll *mongo.Collection) *Handler {
	return &Handler{
		coll: coll,
	}
}
