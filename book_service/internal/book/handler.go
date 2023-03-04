package book

import (
	"context"

	"github.com/Levap123/book_service/proto"
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
}

// type BookServer interface {
//
// 	Delete(context.Context, *DeleteBookRequestResponse) (*DeleteBookRequestResponse, error)
// 	GetAll(*emptypb.Empty, Book_GetAllServer) error
// 	GetByID(context.Context, *GetBookRequset) (*BookInfo, error)
// 	GetByAuthor(context.Context, *GetByAuthorRequest) (*BookInfo, error)
// 	GetByPublisher(context.Context, *GetByPublisherRequest) (*BookInfo, error)
// 	GetByGenre(context.Context, *GetByGenreRequest) (*BookInfo, error)
// 	GetByLanguage(context.Context, *GetByLanguageRequest) (*BookInfo, error)
// 	mustEmbedUnimplementedBookServer()
// }
