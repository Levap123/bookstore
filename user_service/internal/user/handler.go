package user

import (
	"context"
	"fmt"
	"user_service/proto"

	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	proto.UnimplementedUserServer
	logger  *logrus.Logger
	service IUserService
}

func NewUserHandler(service IUserService, logger *logrus.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		logger:  logger,
	}
}

type IUserService interface {
	Create(ctx context.Context, user *CreateUserDTO) (uint64, error)
	GenerateTokens(ctx context.Context, dto *GetUserDTO) (string, string, error)
}

func (uh *UserHandler) SignUp(ctx context.Context, req *proto.SignUpRequest) (*proto.SignUpResponse, error) {
	uh.logger.Debugln("signup user")

	dto := NewCreateUserDTO(req)

	userID, err := uh.service.Create(ctx, dto)
	if err != nil {
		uh.logger.Errorf("error in creating user: %v", err)
		return nil, fmt.Errorf("user handler - signup - %w", err)
	}

	return &proto.SignUpResponse{
		UserID: uint64(userID),
	}, nil
}

func (uh *UserHandler) SignIn(ctx context.Context, req *proto.SignInRequest) (*proto.SignInResponse, error) {
	uh.logger.Debugln("signin user")

	dto := NewGetUserDTO(req)

	accessToken, refreshToken, err := uh.service.GenerateTokens(ctx, dto)
	if err != nil {
		uh.logger.Errorf("error in signin: %v", err)
		return nil, fmt.Errorf("user handler - signin - %w", err)
	}
	return &proto.SignInResponse{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

/* TODO implement:
type UserServer interface {
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	ValidateUser(context.Context, *ValidateRequest) (*ValidateResponse, error)
}
*/
