package user

import (
	"context"
	"errors"
	"fmt"

	"user_service/internal/domain"
	"user_service/internal/validator"
	"user_service/proto"

	"github.com/Levap123/utils/apperror"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	proto.UnimplementedUserServer
	logger    *logrus.Logger
	service   IUserService
	validator *validator.Validator
}

func NewUserHandler(service IUserService, logger *logrus.Logger, validator *validator.Validator) *UserHandler {
	return &UserHandler{
		service:   service,
		logger:    logger,
		validator: validator,
	}
}

type IUserService interface {
	Create(ctx context.Context, user *CreateUserDTO) (uint64, error)
	GenerateTokens(ctx context.Context, dto *GetUserDTO) (string, string, error)
	Validate(ctx context.Context, accessToken string) (int, error)
	GetByID(ctx context.Context, userID uint64) (*User, error)
	UpdateUser(ctx context.Context, dto *UpdateUserDTO) (int, error)
	RefreshTokens(ctx context.Context, accessToken, refreshToken string) (string, string, error)
}

func (uh *UserHandler) SignUp(ctx context.Context, req *proto.SignUpRequest) (*proto.SignUpResponse, error) {
	uh.logger.Debugln("signup user")

	dto := NewCreateUserDTO(req)

	if !uh.validator.IsPasswordLenghtCorrect(dto.Password) {
		return nil, apperror.MakeBadRequestErr(domain.ErrPasswordLengthIncorrect, fmt.Sprintf("password length should be from %d to %d",
			uh.validator.PasswordMin, uh.validator.PasswordMax))
	}

	if !uh.validator.IsUsernameLengthCorrect(dto.Username) {
		return nil, apperror.MakeBadRequestErr(domain.ErrUsernameLengthIncorrect, fmt.Sprintf("username length should be from %d to %d",
			uh.validator.UsernameMin, uh.validator.UsernameMax))
	}

	if !uh.validator.IsEmailCorrect(dto.Email) {
		return nil, apperror.MakeBadRequestErr(domain.ErrIncorrectEmail, domain.ErrIncorrectEmail.Error())
	}

	userID, err := uh.service.Create(ctx, dto)
	if err != nil {
		uh.logger.Errorf("error in creating user: %v", err)

		if errors.Is(err, domain.ErrUnique) {
			return nil, apperror.MakeBadRequestErr(err, domain.ErrUnique.Error())
		}
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

		switch {
		case errors.Is(err, domain.ErrIncorrectPassword):
			return nil, apperror.MakeBadRequestErr(err, domain.ErrIncorrectPassword.Error())
		case errors.Is(err, domain.ErrUserNotFound):
			return nil, apperror.MakeBadRequestErr(err, "check that you print correct email")
		default:
			return nil, fmt.Errorf("user handler - signin - %w", err)
		}
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
		return nil, apperror.MakeUnoauthorizedErr(err)
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
		return nil, apperror.MakeUnoauthorizedErr(err)
	}

	user, err := uh.service.GetByID(ctx, uint64(userID))
	if err != nil {
		uh.logger.Errorf("error in get user by id: %v", err)

		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, apperror.MakeNotFoundErr(err, "user with this id not found")
		}
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

		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, apperror.MakeNotFoundErr(err, "user with this id not found")
		}
		return nil, err
	}
	return &proto.GetResponse{
		UserID:   user.ID,
		Email:    user.Email,
		Username: user.Username,
	}, nil
}

func (uh *UserHandler) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	uh.logger.Debugln("update user credentials")
	dto := NewUpdateUserDTO(req)

	userID, err := uh.service.UpdateUser(ctx, dto)
	if err != nil {
		uh.logger.Errorf("error in updating user: %v", err)

		switch {
		case errors.Is(err, domain.ErrIncorrectPassword):
			return nil, apperror.MakeBadRequestErr(err, domain.ErrIncorrectPassword.Error())
		case errors.Is(err, domain.ErrUnique):
			return nil, apperror.MakeBadRequestErr(err, "this username is busy")
		case errors.Is(err, domain.ErrUserNotFound):
			return nil, apperror.MakeNotFoundErr(err, "user with this id not found")
		default:
			return nil, err
		}
	}

	return &proto.UpdateUserResponse{
		UserID: uint64(userID),
	}, nil
}

func (uh *UserHandler) Refresh(ctx context.Context, req *proto.RefreshRequestResponse) (*proto.RefreshRequestResponse, error) {
	uh.logger.Debugln("refresh access and refresh tokens")

	accessToken, refreshToken, err := uh.service.RefreshTokens(ctx, req.Access, req.Refresh)
	if err != nil {
		uh.logger.Errorf("error in refreshing tokens: %v", err)

		return nil, apperror.MakeUnoauthorizedErr(err)
	}

	return &proto.RefreshRequestResponse{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}
