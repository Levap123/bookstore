package book

import (
	"context"

	"github.com/Levap123/book_service/internal/domain"
)

type BookService struct {
	repo IBookRepo
}

type IBookRepo interface {
	Create(ctx context.Context, book *Book) (string, error)
	GetByID(ctx context.Context, bookID string) (*Book, error)
	GetAll(ctx context.Context) ([]*Book, error)
	GetByAuthor(ctx context.Context, author string) ([]*Book, error)
	GetByPublisher(ctx context.Context, publisher string) ([]*Book, error)
	GetByLanguage(ctx context.Context, language string) ([]*Book, error)
	GetByGenre(ctx context.Context, genre string) ([]*Book, error)
}

func NewBookService(repo IBookRepo) *BookService {
	return &BookService{
		repo: repo,
	}
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
		return nil, err
	}

	if book == nil {
		return nil, domain.ErrBookNotFound
	}

	return book, err
}

func (bs *BookService) GetAll(ctx context.Context) ([]*Book, error) {
	return bs.repo.GetAll(ctx)
}

func (bs *BookService) GetByAuthor(ctx context.Context, author string) ([]*Book, error) {
	return bs.repo.GetByAuthor(ctx, author)
}

func (bs *BookService) GetByPublisher(ctx context.Context, publisher string) ([]*Book, error) {
	return bs.repo.GetByPublisher(ctx, publisher)
}

func (bs *BookService) GetByLanguage(ctx context.Context, language string) ([]*Book, error) {
	return bs.repo.GetByLanguage(ctx, language)
}

func (bs *BookService) GetByGenre(ctx context.Context, genre string) ([]*Book, error) {
	return bs.repo.GetByGenre(ctx, genre)
}
