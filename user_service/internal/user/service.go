package user

import (
	"context"
)

type UserService struct {
	repo UserRepo
}

type UserRepo interface {
	Create(ctx context.Context, user *User) (int, error)
}
