package book

import (
	"context"

	"github.com/Levap123/book_service/internal/domain"
)

type BookService struct {
	repo IBookRepo
}

type IBookRepo interface {
	Delete(ctx context.Context, bookID string) (string, error)
	Create(ctx context.Context, book Book) (string, error)
	GetByID(ctx context.Context, bookID string) (Book, error)
	GetAll(ctx context.Context) ([]Book, error)
	BooksFilter(ctx context.Context, genre, author, language, publisher []string) ([]Book, error)
}

func NewBookService(repo IBookRepo) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (bs *BookService) Create(ctx context.Context, book Book) (string, error) {
	bookID, err := bs.repo.Create(ctx, book)
	if err != nil {
		return "", err
	}
	return bookID, nil
}

func (bs *BookService) GetByID(ctx context.Context, bookID string) (Book, error) {
	book, err := bs.repo.GetByID(ctx, bookID)
	if err != nil {
		return Book{}, err
	}

	if book.ID == "" {
		return Book{}, domain.ErrBookNotFound
	}

	return book, err
}

func (bs *BookService) GetAll(ctx context.Context) ([]Book, error) {
	return bs.repo.GetAll(ctx)
}

func (bs *BookService) Delete(ctx context.Context, bookID string) (string, error) {
	return bs.repo.Delete(ctx, bookID)
}

func (bs *BookService) BooksFilter(ctx context.Context, genre, author, language, publisher []string) ([]Book, error) {
	return bs.repo.BooksFilter(ctx, genre, author, language, publisher)
}
