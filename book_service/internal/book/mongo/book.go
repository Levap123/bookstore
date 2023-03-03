package mongo

import "go.mongodb.org/mongo-driver/mongo"

type BookRepo struct {
	coll *mongo.Collection
}

func NewBookRepo(DB *mongo.Client) *BookRepo {
	return &BookRepo{
		coll: DB.Database("bookstore").Collection("books"),
	}
}
