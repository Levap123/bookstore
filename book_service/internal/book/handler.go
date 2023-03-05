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
	Delete(ctx context.Context, bookID string) (string, error)
	Create(ctx context.Context, book *Book) (string, error)
	GetByID(ctx context.Context, bookID string) (*Book, error)
	GetAll(ctx context.Context) ([]*Book, error)
	GetByAuthor(ctx context.Context, author string) ([]*Book, error)
	GetByPublisher(ctx context.Context, author string) ([]*Book, error)
	GetByLanguage(ctx context.Context, author string) ([]*Book, error)
	GetByGenre(ctx context.Context, author string) ([]*Book, error)
}

func NewBookHandler(service IBookService, log *logrus.Logger) *BookHandler {
	return &BookHandler{
		service: service,
		log:     log,
	}
}

func (h *BookHandler) Create(ctx context.Context, req *proto.BookInfo) (*proto.CreateBookResponse, error) {
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
		h.log.Errorf("error in getting all books: %v", err)

		if errors.Is(err, domain.ErrBookNotFound) {
			return nil, status.Errorf(codes.NotFound, domain.ErrBookNotFound.Error())
		}
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

func (h *BookHandler) GetByAuthor(ctx context.Context, req *proto.GetByAuthorRequest) (*proto.BookInfoArray, error) {
	books, err := h.service.GetByAuthor(ctx, req.Author)
	if err != nil {
		h.log.Errorf("error in getting all books by author: %v", err)

		if errors.Is(err, domain.ErrBookNotFound) {
			return nil, status.Errorf(codes.NotFound, domain.ErrBookNotFound.Error())
		}
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

func (h *BookHandler) GetByPublisher(ctx context.Context, req *proto.GetByPublisherRequest) (*proto.BookInfoArray, error) {
	books, err := h.service.GetByPublisher(ctx, req.Publisher)
	if err != nil {
		h.log.Errorf("error in getting all books by publisher: %v", err)

		if errors.Is(err, domain.ErrBookNotFound) {
			return nil, status.Errorf(codes.NotFound, domain.ErrBookNotFound.Error())
		}
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

func (h *BookHandler) GetByGenre(ctx context.Context, req *proto.GetByGenreRequest) (*proto.BookInfoArray, error) {
	books, err := h.service.GetByGenre(ctx, req.Genre)
	if err != nil {
		h.log.Errorf("error in getting all books by genre: %v", err)

		if errors.Is(err, domain.ErrBookNotFound) {
			return nil, status.Errorf(codes.NotFound, domain.ErrBookNotFound.Error())
		}
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

func (h *BookHandler) GetByLanguage(ctx context.Context, req *proto.GetByLanguageRequest) (*proto.BookInfoArray, error) {
	books, err := h.service.GetByLanguage(ctx, req.Language)
	if err != nil {
		h.log.Errorf("error in getting all books by language: %v", err)

		if errors.Is(err, domain.ErrBookNotFound) {
			return nil, status.Errorf(codes.NotFound, domain.ErrBookNotFound.Error())
		}
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

func (h *BookHandler) Delete(ctx context.Context, req *proto.DeleteBookRequestResponse) (*proto.DeleteBookRequestResponse, error) {
	bookID, err := h.service.Delete(ctx, req.BookID)
	if err != nil {
		h.log.Errorf("error in deleting book: %v", err)
		if errors.Is(err, domain.ErrBookNotFound) {
			return nil, status.Errorf(codes.NotFound, "book did not deleted because book with this id not found")
		}
		return nil, err
	}
	return &proto.DeleteBookRequestResponse{
		BookID: bookID,
	}, nil
}

/* TODO: implement
type BookServer interface {
	mustEmbedUnimplementedBookServer()
}
*/
