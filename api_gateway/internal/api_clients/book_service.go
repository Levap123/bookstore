package apiclients

import (
	"context"

	"github.com/Levap123/api_gateway/internal/dto"
	"github.com/Levap123/api_gateway/internal/entity"
	"github.com/Levap123/api_gateway/proto"
	"github.com/Levap123/utils/apperror"
	"github.com/golang/protobuf/ptypes/empty"
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
	resp, err := bc.cl.GetAll(ctx, &empty.Empty{})
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

func (bc *BookClient) GetByFiltering(ctx context.Context, params map[string][]string) ([]entity.Book, error) {
	filter := &proto.Filter{
		Author:    params["author"],
		Genre:     params["genre"],
		Language:  params["language"],
		Publsiher: params["publisher"],
	}

	resp, err := bc.cl.GetWithFilter(ctx, filter)
	if err != nil {
		return nil, err
	}

	bookArr := make([]entity.Book, 0, len(resp.Arr))

	for _, bookReq := range resp.Arr {
		bookArr = append(bookArr, entity.FromBookRequestToBook(bookReq))
	}

	return bookArr, nil
}

