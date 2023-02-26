package user

import (
	"context"
	"time"
	"user_service/proto"
)

type UserHandler struct {
	proto.UnimplementedUserServer
	service IUserService
}

func NewUserHandler(service IUserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

type IUserService interface {
	Create(ctx context.Context, user *CreateUserDTO) (int, error)
}

func (uh *UserHandler) SignUp(context.Context, *proto.SignUpRequest) (*proto.SignUpResponse, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

}

/* TODO implement:
type UserServer interface {
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	ValidateUser(context.Context, *ValidateRequest) (*ValidateResponse, error)
	mustEmbedUnimplementedUserServer()
}
*/
