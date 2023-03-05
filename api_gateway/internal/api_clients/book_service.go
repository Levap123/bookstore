package apiclients

import (
	"context"

	"github.com/Levap123/api_gateway/internal/dto"
	"github.com/Levap123/api_gateway/internal/entity"
	"github.com/Levap123/api_gateway/proto"
	"github.com/Levap123/utils/apperror"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type BookClient struct {
	cl  proto.BookClient
	log *logrus.Logger
}

func InitBookClient(conn *grpc.ClientConn, log *logrus.Logger) *BookClient {
	cl := proto.NewBookClient(conn)
	return &BookClient{
		cl:  cl,
		log: log,
	}
}

func (bc *BookClient) Create(ctx context.Context, createBookDTO dto.CreateBookDTO) (string, error) {
	bookRequest := dto.FromDtoToRequest(createBookDTO)

	resp, err := bc.cl.Create(ctx, bookRequest)
	if err != nil {
		bc.log.Errorf("error from book service: %v", err)

		status, ok := status.FromError(err)
		if !ok {
			return "", err
		}

		statusCode := gRPCToHTTP(status.Code())
		if statusCode == -1 {
			return "", err
		}

		return "", apperror.NewError(status.Err(), status.Message(), statusCode)
	}
	return resp.BookID, nil
}

func (bc *BookClient) GetByID(ctx context.Context, bookID string) (entity.Book, error) {
	req := &proto.GetBookRequset{
		BookID: bookID,
	}

	resp, err := bc.cl.GetByID(ctx, req)
	if err != nil {
		bc.log.Errorf("error from book service: %v", err)

		status, ok := status.FromError(err)
		if !ok {
			return entity.Book{}, err
		}

		statusCode := gRPCToHTTP(status.Code())
		if statusCode == -1 {
			return entity.Book{}, err
		}

		return entity.Book{}, apperror.NewError(status.Err(), status.Message(), statusCode)
	}

	return entity.FromBookRequestToBook(resp), nil
}

func (bc *BookClient) Delete(ctx context.Context, bookID string) (string, error) {
	req := &proto.DeleteBookRequestResponse{
		BookID: bookID,
	}

	resp, err := bc.cl.Delete(ctx, req)
	if err != nil {
		bc.log.Errorf("error from book service: %v", err)

		status, ok := status.FromError(err)
		if !ok {
			return "", err
		}

		statusCode := gRPCToHTTP(status.Code())
		if statusCode == -1 {
			return "", err
		}

		return "", apperror.NewError(status.Err(), status.Message(), statusCode)
	}

	return resp.BookID, err
}

func (bc *BookClient) GetAll(ctx context.Context) ([]entity.Book, error) {
	resp, err := bc.cl.GetAll(ctx, nil)
	if err != nil {
		bc.log.Errorf("error from book service: %v", err)

		status, ok := status.FromError(err)
		if !ok {
			return nil, err
		}

		statusCode := gRPCToHTTP(status.Code())
		if statusCode == -1 {
			return nil, err
		}

		return nil, apperror.NewError(status.Err(), status.Message(), statusCode)
	}

	bookArr := make([]entity.Book, 0, len(resp.Arr))

	for _, bookReq := range resp.Arr {
		bookArr = append(bookArr, entity.FromBookRequestToBook(bookReq))
	}

	return bookArr, nil
}

func (bc *BookClient) GetAllByAuthor(ctx context.Context, author string) ([]entity.Book, error) {
	req := &proto.GetByAuthorRequest{
		Author: author,
	}

	resp, err := bc.cl.GetByAuthor(ctx, req)
	if err != nil {
		bc.log.Errorf("error from book service: %v", err)

		status, ok := status.FromError(err)
		if !ok {
			return nil, err
		}

		statusCode := gRPCToHTTP(status.Code())
		if statusCode == -1 {
			return nil, err
		}

		return nil, apperror.NewError(status.Err(), status.Message(), statusCode)
	}

	bookArr := make([]entity.Book, 0, len(resp.Arr))

	for _, bookReq := range resp.Arr {
		bookArr = append(bookArr, entity.FromBookRequestToBook(bookReq))
	}

	return bookArr, nil
}

func (bc *BookClient) GetAllByLanguage(ctx context.Context, language string) ([]entity.Book, error) {
	req := &proto.GetByLanguageRequest{
		Language: language,
	}

	resp, err := bc.cl.GetByLanguage(ctx, req)
	if err != nil {
		bc.log.Errorf("error from book service: %v", err)

		status, ok := status.FromError(err)
		if !ok {
			return nil, err
		}

		statusCode := gRPCToHTTP(status.Code())
		if statusCode == -1 {
			return nil, err
		}

		return nil, apperror.NewError(status.Err(), status.Message(), statusCode)
	}

	bookArr := make([]entity.Book, 0, len(resp.Arr))

	for _, bookReq := range resp.Arr {
		bookArr = append(bookArr, entity.FromBookRequestToBook(bookReq))
	}

	return bookArr, nil
}

func (bc *BookClient) GetAllByGenre(ctx context.Context, genre string) ([]entity.Book, error) {
	req := &proto.GetByGenreRequest{
		Genre: genre,
	}

	resp, err := bc.cl.GetByGenre(ctx, req)
	if err != nil {
		bc.log.Errorf("error from book service: %v", err)

		status, ok := status.FromError(err)
		if !ok {
			return nil, err
		}

		statusCode := gRPCToHTTP(status.Code())
		if statusCode == -1 {
			return nil, err
		}

		return nil, apperror.NewError(status.Err(), status.Message(), statusCode)
	}

	bookArr := make([]entity.Book, 0, len(resp.Arr))

	for _, bookReq := range resp.Arr {
		bookArr = append(bookArr, entity.FromBookRequestToBook(bookReq))
	}

	return bookArr, nil
}

func (bc *BookClient) GetAllByPublisher(ctx context.Context, publisher string) ([]entity.Book, error) {
	req := &proto.GetByPublisherRequest{
		Publisher: publisher,
	}

	resp, err := bc.cl.GetByPublisher(ctx, req)
	if err != nil {
		bc.log.Errorf("error from book service: %v", err)

		status, ok := status.FromError(err)
		if !ok {
			return nil, err
		}

		statusCode := gRPCToHTTP(status.Code())
		if statusCode == -1 {
			return nil, err
		}

		return nil, apperror.NewError(status.Err(), status.Message(), statusCode)
	}

	bookArr := make([]entity.Book, 0, len(resp.Arr))

	for _, bookReq := range resp.Arr {
		bookArr = append(bookArr, entity.FromBookRequestToBook(bookReq))
	}

	return bookArr, nil
}
