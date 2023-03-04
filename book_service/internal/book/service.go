package book

import (
	"context"
	"errors"

	"github.com/Levap123/book_service/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookService struct {
	repo IBookRepo
}

type IBookRepo interface {
	Create(ctx context.Context, book *Book) (string, error)
	GetByID(ctx context.Context, bookID string) (*Book, error)
}

func (bs *BookService) Create(ctx context.Context, book *Book) (string, error) {
	bookID, err := bs.repo.Create(ctx, book)
	if err != nil {
		return "", err
	}
	return bookID, nil
}

func (bs *BookService) GetByID(ctx context.Context, bookID string) (*Book, error) {
	book, err := bs.repo.GetByID(ctx, bookID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrBookNotFound
		}
		return nil, err
	}
	if book == nil {
		return nil, domain.ErrBookNotFound
	}
	return book, err
}
