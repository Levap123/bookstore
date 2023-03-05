package mongo

import (
	"context"
	"errors"
	"fmt"

	"github.com/Levap123/book_service/internal/book"
	"github.com/Levap123/book_service/internal/domain"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepo struct {
	coll *mongo.Collection
	log  *logrus.Logger
}

func NewBookRepo(DB *mongo.Client, log *logrus.Logger) *BookRepo {
	return &BookRepo{
		coll: DB.Database("bookstore").Collection("books"),
		log:  log,
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

func (br *BookRepo) GetByID(ctx context.Context, bookID string) (*book.Book, error) {
	objectID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return nil, fmt.Errorf("book repo - get object ID from hex - %w", err)
	}

	filter := bson.D{
		{"_id", objectID},
	}

	var book book.Book
	if err := br.coll.FindOne(ctx, filter).Decode(&book); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrBookNotFound
		}
		return nil, fmt.Errorf("book repo - get one - %w", err)
	}

	return &book, err
}

func (br *BookRepo) GetAll(ctx context.Context) ([]*book.Book, error) {

	cur, err := br.coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("book repo - get all - %w", err)
	}

	books := make([]*book.Book, 0, 10)

	for cur.Next(ctx) {
		var buffer book.Book
		if err := cur.Decode(&buffer); err != nil {
			return nil, fmt.Errorf("book repo - get all - %w", err)
		}

		books = append(books, &buffer)
	}

	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("book repo - get all - %w", err)
	}

	if len(books) == 0 {
		return nil, domain.ErrBookNotFound
	}

	return books, nil
}

func (br *BookRepo) GetByAuthor(ctx context.Context, author string) ([]*book.Book, error) {

	filter := bson.D{
		{"author", author},
	}
	cur, err := br.coll.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("book repo - get all - %w", err)
	}

	books := make([]*book.Book, 0, 10)

	for cur.Next(ctx) {
		var buffer book.Book
		if err := cur.Decode(&buffer); err != nil {
			return nil, fmt.Errorf("book repo - get all - %w", err)
		}

		books = append(books, &buffer)
	}

	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("book repo - get all - %w", err)
	}

	if len(books) == 0 {
		return nil, domain.ErrBookNotFound
	}

	return books, nil
}

func (br *BookRepo) GetByPublisher(ctx context.Context, publisher string) ([]*book.Book, error) {

	filter := bson.D{
		{"publisher", publisher},
	}
	cur, err := br.coll.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("book repo - get all - %w", err)
	}

	books := make([]*book.Book, 0, 10)

	for cur.Next(ctx) {
		var buffer book.Book
		if err := cur.Decode(&buffer); err != nil {
			return nil, fmt.Errorf("book repo - get all - %w", err)
		}

		books = append(books, &buffer)
	}

	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("book repo - get all - %w", err)
	}

	if len(books) == 0 {
		return nil, domain.ErrBookNotFound
	}

	return books, nil
}

func (br *BookRepo) GetByGenre(ctx context.Context, genre string) ([]*book.Book, error) {

	filter := bson.D{
		{"genre", genre},
	}
	cur, err := br.coll.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("book repo - get all - %w", err)
	}

	books := make([]*book.Book, 0, 10)

	for cur.Next(ctx) {
		var buffer book.Book
		if err := cur.Decode(&buffer); err != nil {
			return nil, fmt.Errorf("book repo - get all - %w", err)
		}

		books = append(books, &buffer)
	}

	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("book repo - get all - %w", err)
	}

	if len(books) == 0 {
		return nil, domain.ErrBookNotFound
	}

	return books, nil
}

func (br *BookRepo) GetByLanguage(ctx context.Context, language string) ([]*book.Book, error) {

	filter := bson.D{
		{"language", language},
	}
	cur, err := br.coll.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("book repo - get all - %w", err)
	}

	books := make([]*book.Book, 0, 10)

	for cur.Next(ctx) {
		var buffer book.Book
		if err := cur.Decode(&buffer); err != nil {
			return nil, fmt.Errorf("book repo - get all - %w", err)
		}

		books = append(books, &buffer)
	}

	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("book repo - get all - %w", err)
	}

	if len(books) == 0 {
		return nil, domain.ErrBookNotFound
	}

	return books, nil
}

func (br *BookRepo) Delete(ctx context.Context, bookID string) (string, error) {
	objectID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return "", fmt.Errorf("book repo - get object ID from hex - %w", domain.ErrBookNotFound)
	}

	filter := bson.D{
		{"_id", objectID},
	}

	res, err := br.coll.DeleteOne(ctx, filter)
	if err != nil {
		return "", fmt.Errorf("book repo - delete by id - %w", err)
	}
	if res.DeletedCount != 1 {
		return "", domain.ErrBookNotFound
	}

	return bookID, nil
}
