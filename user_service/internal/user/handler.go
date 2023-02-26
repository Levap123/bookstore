package user

import (
	"context"
	"fmt"
	"time"
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
}

func (uh *UserHandler) SignUp(ctx context.Context, req *proto.SignUpRequest) (*proto.SignUpResponse, error) {
	uh.logger.Debugln("signup user")
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	user := NewCreateUserDTO(req)
	userID, err := uh.service.Create(ctx, user)
	if err != nil {
		uh.logger.Errorf("error in creating user: %w\n", err)
		return nil, fmt.Errorf("user handler - signup - %w", err)
	}
	return &proto.SignUpResponse{
		UserID: uint64(userID),
	}, nil
}

/* TODO implement:
type UserServer interface {
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	ValidateUser(context.Context, *ValidateRequest) (*ValidateResponse, error)
	mustEmbedUnimplementedUserServer()
}
*/
