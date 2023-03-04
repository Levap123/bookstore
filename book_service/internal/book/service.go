package book

import "context"

type BookService struct {
	repo BookRepo
}

type BookRepo interface {
	Create(ctx context.Context, book *Book) (string, error)
	Get(ctx context.Context, bookID string) (*Book, error)
}

func (bs *BookService) Create(ctx context.Context, book *Book) (string, error) {
	return bs.repo.Create(ctx, book)
}

func (bs *BookService) Get(ctx context.Context, bookID string) (*Book, error) {
	return bs.repo.Get(ctx, bookID)
}
