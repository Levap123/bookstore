package apiclients

import (
	"api-gateway/internal/dto"
	"api-gateway/proto"
	"context"

	"google.golang.org/grpc"
)

type UserClient struct {
	cl proto.UserClient
}

func InitUserClient(conn *grpc.ClientConn) *UserClient {
	cl := proto.NewUserClient(conn)
	return &UserClient{
		cl: cl,
	}
}

func (uc *UserClient) SignUp(ctx context.Context, dto *dto.SignUpDTO) (uint64, error) {
	request := &proto.SignUpRequest{
		Username: dto.Username,
		Email:    dto.Email,
		Password: dto.Password,
	}

	response, err := uc.cl.SignUp(ctx, request)
	if err != nil {
		return 0, err
	}

	return response.UserID, nil
}
