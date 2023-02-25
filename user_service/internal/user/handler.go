package user

import (
	"context"
	"user_service/proto"
)

type UserHandler struct {
	proto.UnimplementedUserServer
}

func (uh *UserHandler) SignUp(context.Context, *proto.SignUpRequest) (*proto.SignUpResponse, error) {
}

/* TODO implement:
type UserServer interface {
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	ValidateUser(context.Context, *ValidateRequest) (*ValidateResponse, error)
	mustEmbedUnimplementedUserServer()
}
*/
