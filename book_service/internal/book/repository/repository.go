package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Levap123/book_service/internal/book"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type Repo struct {
	repo  IBookRepo
	cache *redis.Client
	log   *logrus.Logger
}

func NewRepo(repo IBookRepo, cache *redis.Client, log *logrus.Logger) *Repo {
	return &Repo{
		repo:  repo,
		cache: cache,
		log:   log,
	}
}

type IBookRepo interface {
	Delete(ctx context.Context, bookID string) (string, error)
	Create(ctx context.Context, book book.Book) (string, error)
	GetByID(ctx context.Context, bookID string) (book.Book, error)
	GetAll(ctx context.Context) ([]book.Book, error)
	BooksFilter(ctx context.Context, genre, author, language, publisher []string) ([]book.Book, error)
}

const allBooksKey = "books:all"

func (r *Repo) Delete(ctx context.Context, bookID string) (string, error) {
	return r.repo.Delete(ctx, bookID)
}

func (r *Repo) Create(ctx context.Context, book book.Book) (string, error) {
	return r.repo.Create(ctx, book)
}

func (r *Repo) GetByID(ctx context.Context, bookID string) (book.Book, error) {
	return r.repo.GetByID(ctx, bookID)
}

func (r *Repo) BooksFilter(ctx context.Context, genre, author, language, publisher []string) ([]book.Book, error) {
	return r.repo.BooksFilter(ctx, genre, author, language, publisher)
}

func (r *Repo) GetAll(ctx context.Context) ([]book.Book, error) {
	getAllBytes := r.getFromRedis(ctx, allBooksKey)

	if len(getAllBytes) != 0 {
		var books []book.Book
		if err := json.Unmarshal(getAllBytes, &books); err == nil {
			r.log.Debug("get all books - getting from redis")
			return books, nil
		} else {
			r.log.Error("repo - error in unmarshalling request - %v", err)
		}
	}

	books, err := r.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	booksBytes, err := json.Marshal(books)
	if err != nil {
		r.log.Errorf("repo - error in marshalling request - %w", err)
	} else {
		status := r.cache.Set(ctx, allBooksKey, booksBytes, time.Minute*5)
		if status.Err() != nil {
			r.log.Errorf("repo - error in setting request to redis - %w", err)
		}
	}
	return books, nil
}

func (r *Repo) getFromRedis(ctx context.Context, key string) []byte {
	cache, err := r.cache.Get(ctx, key).Bytes()
	if err != nil {
		return []byte{}
	}
	return cache
}
