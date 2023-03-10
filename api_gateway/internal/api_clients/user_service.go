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

type UserClient struct {
	cl  proto.UserClient
	log *logrus.Logger
}

func InitUserClient(conn *grpc.ClientConn, log *logrus.Logger) *UserClient {
	cl := proto.NewUserClient(conn)
	return &UserClient{
		cl:  cl,
		log: log,
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
		uc.log.Errorf("error from user service: %v", err)

		status, ok := status.FromError(err)
		if !ok {
			return 0, err
		}

		statusCode := gRPCToHTTP(status.Code())
		if statusCode == -1 {
			return 0, err
		}

		return 0, apperror.NewError(status.Err(), status.Message(), statusCode)
	}

	return response.UserID, nil
}

func (uc *UserClient) SignIn(ctx context.Context, dto *dto.SignInDTO) (*entity.Tokens, error) {
	request := &proto.SignInRequest{
		Email:    dto.Email,
		Password: dto.Password,
	}

	response, err := uc.cl.SignIn(ctx, request)
	if err != nil {
		uc.log.Errorf("error from user service: %v", err)

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

	return &entity.Tokens{
		Access:  response.Access,
		Refresh: response.Refresh,
	}, nil
}

func (uc *UserClient) Validate(ctx context.Context, accessToken string) (uint64, error) {
	request := &proto.ValidateRequest{
		Access: accessToken,
	}

	response, err := uc.cl.ValidateUser(ctx, request)
	if err != nil {
		uc.log.Errorf("error from user service: %v", err)

		status, ok := status.FromError(err)
		if !ok {
			return 0, err
		}

		statusCode := gRPCToHTTP(status.Code())
		if statusCode == -1 {
			return 0, err
		}

		return 0, apperror.NewError(status.Err(), status.Message(), statusCode)
	}

	return response.UserID, nil
}

func (uc *UserClient) GetMe(ctx context.Context, accessToken string) (*entity.User, error) {
	request := &proto.ValidateRequest{
		Access: accessToken,
	}

	response, err := uc.cl.GetMe(ctx, request)
	if err != nil {
		uc.log.Errorf("error from user service: %v", err)

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

	return &entity.User{
		ID:       response.UserID,
		Email:    response.Email,
		Username: response.Username,
	}, nil
}

func (uc *UserClient) GetByID(ctx context.Context, userID uint64) (*entity.User, error) {
	request := &proto.GetByIDRequest{
		UserID: userID,
	}

	response, err := uc.cl.GetById(ctx, request)
	if err != nil {
		uc.log.Errorf("error from user service: %v", err)

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

	return &entity.User{
		ID:       response.UserID,
		Email:    response.Email,
		Username: response.Username,
	}, nil
}

func (uc *UserClient) Update(ctx context.Context, dto *dto.UpdateUserDTO) (uint64, error) {
	request := &proto.UpdateUserRequest{
		Email:       dto.Email,
		Username:    dto.Username,
		NewPassword: dto.NewPassword,
		OldPassword: dto.OldPassword,
	}

	response, err := uc.cl.UpdateUser(ctx, request)
	if err != nil {
		uc.log.Errorf("error from user service: %v", err)

		status, ok := status.FromError(err)
		if !ok {
			return 0, err
		}

		statusCode := gRPCToHTTP(status.Code())
		if statusCode == -1 {
			return 0, err
		}

		return 0, apperror.NewError(status.Err(), status.Message(), statusCode)
	}

	return response.UserID, nil
}

func (uc *UserClient) Rerfresh(ctx context.Context, dtoReq *dto.RefreshDTO) (*dto.RefreshDTO, error) {
	request := &proto.RefreshRequestResponse{
		Access:  dtoReq.AccessToken,
		Refresh: dtoReq.RefreshToken,
	}

	response, err := uc.cl.Refresh(ctx, request)
	if err != nil {
		uc.log.Errorf("error from user service: %v", err)

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

	return &dto.RefreshDTO{
		AccessToken:  response.Access,
		RefreshToken: response.Refresh,
	}, nil
}
