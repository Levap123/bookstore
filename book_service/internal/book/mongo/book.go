package mongo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Levap123/book_service/internal/book"
	"github.com/Levap123/book_service/internal/domain"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepo struct {
	coll  *mongo.Collection
	cache *redis.Client
	log   *logrus.Logger
}

func NewBookRepo(DB *mongo.Client, redisCleint *redis.Client, log *logrus.Logger) *BookRepo {
	return &BookRepo{
		coll:  DB.Database("bookstore").Collection("books"),
		cache: redisCleint,
		log:   log,
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
	getAllBytes := br.getFromRedis(ctx, "all")

	if len(getAllBytes) != 0 {
		var books []*book.Book
		if err := json.Unmarshal(getAllBytes, &books); err == nil {
			br.log.Debug("getting from redis")
			return books, nil
		}
	}

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

	booksBytes, err := json.Marshal(books)
	if err != nil {
		br.log.Errorf("repo - error in marshalling request - %w", err)
	} else {
		br.cache.Set(ctx, "all", booksBytes, time.Minute*5)
	}

	return books, nil
}

func (br *BookRepo) GetByAuthor(ctx context.Context, author string) ([]*book.Book, error) {
	getAllBytes := br.getFromRedis(ctx, "author")

	if len(getAllBytes) != 0 {
		var books []*book.Book
		if err := json.Unmarshal(getAllBytes, &books); err == nil {
			br.log.Debug("getting from redis")
			return books, nil
		}
	}

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

	booksBytes, err := json.Marshal(books)
	if err != nil {
		br.log.Errorf("repo - error in marshalling request - %w", err)
	} else {
		br.cache.Set(ctx, "author", booksBytes, time.Minute*5)
	}

	return books, nil
}

func (br *BookRepo) GetByPublisher(ctx context.Context, publisher string) ([]*book.Book, error) {
	getAllBytes := br.getFromRedis(ctx, "publisher")

	if len(getAllBytes) != 0 {
		var books []*book.Book
		if err := json.Unmarshal(getAllBytes, &books); err == nil {
			br.log.Debug("getting from redis")
			return books, nil
		}
	}

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

	booksBytes, err := json.Marshal(books)
	if err != nil {
		br.log.Errorf("repo - error in marshalling request - %w", err)
	} else {
		br.cache.Set(ctx, "publisher", booksBytes, time.Minute*5)
	}

	return books, nil
}

func (br *BookRepo) GetByGenre(ctx context.Context, genre string) ([]*book.Book, error) {
	getAllBytes := br.getFromRedis(ctx, "genre")

	if len(getAllBytes) != 0 {
		var books []*book.Book
		if err := json.Unmarshal(getAllBytes, &books); err == nil {
			br.log.Debug("getting from redis")
			return books, nil
		}
	}

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

	booksBytes, err := json.Marshal(books)
	if err != nil {
		br.log.Errorf("repo - error in marshalling request - %w", err)
	} else {
		br.cache.Set(ctx, "genre", booksBytes, time.Minute*5)
	}

	return books, nil
}

func (br *BookRepo) GetByLanguage(ctx context.Context, language string) ([]*book.Book, error) {
	getAllBytes := br.getFromRedis(ctx, "language")

	if len(getAllBytes) != 0 {
		var books []*book.Book
		if err := json.Unmarshal(getAllBytes, &books); err == nil {
			br.log.Debug("getting from redis")
			return books, nil
		}
	}

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

	booksBytes, err := json.Marshal(books)
	if err != nil {
		br.log.Errorf("repo - error in marshalling request - %w", err)
	} else {
		br.cache.Set(ctx, "language", booksBytes, time.Minute*5)
	}

	return books, nil
}

func (br *BookRepo) getFromRedis(ctx context.Context, key string) []byte {
	cache, err := br.cache.Get(ctx, key).Bytes()
	if err != nil {
		return []byte{}
	}
	return cache
}
