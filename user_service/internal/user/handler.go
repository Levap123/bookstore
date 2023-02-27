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
	Validate(ctx context.Context, accessToken string) (int, error)
	GetByID(ctx context.Context, userID uint64) (*User, error)
	UpdateUser(ctx context.Context, dto *UpdateUserDTO) (int, error)
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

func (uh *UserHandler) ValidateUser(ctx context.Context, req *proto.ValidateRequest) (*proto.ValidateResponse, error) {
	uh.logger.Debugln("valivate user access token")

	userID, err := uh.service.Validate(ctx, req.Access)
	if err != nil {
		uh.logger.Errorf("error in parse user token: %v", err)
		return nil, err
	}
	return &proto.ValidateResponse{
		UserID: uint64(userID),
	}, nil
}

func (uh *UserHandler) GetMe(ctx context.Context, req *proto.ValidateRequest) (*proto.GetResponse, error) {
	uh.logger.Debugln("get user info by access token")

	userID, err := uh.service.Validate(ctx, req.Access)
	if err != nil {
		uh.logger.Errorf("error in parse user token: %v", err)
		return nil, err
	}

	user, err := uh.service.GetByID(ctx, uint64(userID))
	if err != nil {
		uh.logger.Errorf("error in get user by id: %v", err)
		return nil, err
	}
	return &proto.GetResponse{
		UserID:   user.ID,
		Email:    user.Email,
		Username: user.Username,
	}, nil
}

func (uh *UserHandler) GetById(ctx context.Context, req *proto.GetByIDRequest) (*proto.GetResponse, error) {
	uh.logger.Debugln("get user by id")

	user, err := uh.service.GetByID(ctx, req.UserID)
	if err != nil {
		uh.logger.Errorf("error in get user by id: %v", err)
		return nil, err
	}
	return &proto.GetResponse{
		UserID:   user.ID,
		Email:    user.Email,
		Username: user.Username,
	}, nil
}

/* TODO implement:
type UserServer interface {
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	mustEmbedUnimplementedUserServer()
}
*/

func (uh *UserHandler) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	uh.logger.Debugln("update user credentials")
	dto := NewUpdateUserDTO(req)

	userID, err := uh.service.UpdateUser(ctx, dto)
	if err != nil {
		uh.logger.Errorf("error in updating user: %v", err)
		return nil, err
	}

	return &proto.UpdateUserResponse{
		UserID: uint64(userID),
	}, nil
}
