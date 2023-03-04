package book

import (
	"context"
	"errors"

	"github.com/Levap123/book_service/internal/domain"
	"github.com/Levap123/book_service/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BookHandler struct {
	proto.UnimplementedBookServer
	service IBookService
	log     *logrus.Logger
}

type IBookService interface {
	Create(ctx context.Context, book *Book) (string, error)
	GetByID(ctx context.Context, bookID string) (*Book, error)
	GetAll(ctx context.Context) ([]*Book, error)
}

func NewBookHandler(service IBookService, log *logrus.Logger) *BookHandler {
	return &BookHandler{
		service: service,
		log:     log,
	}
}

func (h *BookHandler) Create(ctx context.Context, req *proto.CreateBookRequest) (*proto.CreateBookResponse, error) {
	book := NewBookFromCreateBookRequest(req)
	bookID, err := h.service.Create(ctx, book)
	if err != nil {
		h.log.Errorf("error in creating book: %v", err)
		return nil, err
	}
	return &proto.CreateBookResponse{
		BookID: bookID,
	}, nil
}

func (h *BookHandler) GetByID(ctx context.Context, req *proto.GetBookRequset) (*proto.BookInfo, error) {
	book, err := h.service.GetByID(ctx, req.BookID)
	if err != nil {
		h.log.Errorf("error in getting book by ID: %v", err)
		if errors.Is(err, domain.ErrBookNotFound) {
			return nil, status.Errorf(codes.NotFound, "book with this ID not found")
		}
	}

	if book == nil {
		return nil, status.Errorf(codes.NotFound, "book with this ID not found")
	}
	return NewBookResponseFromBook(book), nil
}

func (h *BookHandler) GetAll(ctx context.Context, req *emptypb.Empty) (*proto.BookInfoArray, error) {
	books, err := h.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	requestArray := make([]*proto.BookInfo, 0, len(books))

	for _, book := range books {
		requestArray = append(requestArray, NewBookResponseFromBook(book))
	}

	return &proto.BookInfoArray{
		Arr: requestArray,
	}, nil
}

/* TODO: implement
type BookServer interface {
	Delete(context.Context, *DeleteBookRequestResponse) (*DeleteBookRequestResponse, error)
	GetAll(context.Context, *emptypb.Empty) (*BookInfoArray, error)
	GetByAuthor(context.Context, *GetByAuthorRequest) (*BookInfoArray, error)
	GetByPublisher(context.Context, *GetByPublisherRequest) (*BookInfoArray, error)
	GetByGenre(context.Context, *GetByGenreRequest) (*BookInfoArray, error)
	GetByLanguage(context.Context, *GetByLanguageRequest) (*BookInfoArray, error)
	mustEmbedUnimplementedBookServer()
}
*/
