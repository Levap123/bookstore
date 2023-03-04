package book

import (
	"context"
	"errors"

	"github.com/Levap123/book_service/internal/domain"
	"github.com/Levap123/book_service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookHandler struct {
	proto.UnimplementedBookServer
	service BookService
}

type IBookService interface {
	Create(ctx context.Context, book *Book) (string, error)
	GetByID(ctx context.Context, bookID string) (*Book, error)
}

func (bh *BookHandler) Create(ctx context.Context, req *proto.CreateBookRequest) (*proto.CreateBookResponse, error) {
	book := NewBookFromCreateBookRequest(req)
	bookID, err := bh.service.Create(ctx, book)
	if err != nil {
		return nil, err
	}
	return &proto.CreateBookResponse{
		BookID: bookID,
	}, nil
}

func (bh *BookHandler) GetByID(ctx context.Context, req *proto.GetBookRequset) (*proto.BookInfo, error) {
	book, err := bh.service.GetByID(ctx, req.BookID)
	if err != nil {
		if errors.Is(err, domain.ErrBookNotFound) {
			return nil, status.Errorf(codes.NotFound, "book with this ID not found")
		}
	}
	return NewBookResponseFromBook(book), nil
}

// type BookServer interface {
//
// 	Delete(context.Context, *DeleteBookRequestResponse) (*DeleteBookRequestResponse, error)
// 	GetAll(*emptypb.Empty, Book_GetAllServer) error
//
// 	GetByAuthor(context.Context, *GetByAuthorRequest) (*BookInfo, error)
// 	GetByPublisher(context.Context, *GetByPublisherRequest) (*BookInfo, error)
// 	GetByGenre(context.Context, *GetByGenreRequest) (*BookInfo, error)
// 	GetByLanguage(context.Context, *GetByLanguageRequest) (*BookInfo, error)
// 	mustEmbedUnimplementedBookServer()
// }
