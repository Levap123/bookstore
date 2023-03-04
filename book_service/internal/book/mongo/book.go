package mongo

import (
	"context"
	"fmt"

	"github.com/Levap123/book_service/internal/book"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepo struct {
	coll *mongo.Collection
}

func NewBookRepo(DB *mongo.Client) *BookRepo {
	return &BookRepo{
		coll: DB.Database("bookstore").Collection("books"),
	}
}

func (br *BookRepo) Create(ctx context.Context, book *book.Book) (string, error) {
	res, err := br.coll.InsertOne(ctx, book)
	if err != nil {
		return "", fmt.Errorf("book repo - create - %w", err)
	}

	ID := res.InsertedID.(primitive.ObjectID).Hex()
	return ID, nil
}

func (br *BookRepo) Get(ctx context.Context, bookID string) (*book.Book, error){
	
}
