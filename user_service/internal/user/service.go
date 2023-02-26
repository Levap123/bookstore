package user

import (
	"context"
)

type UserService struct {
	repo IUserRepo
}

func NewUserService(repo IUserRepo) *UserService {
	return &UserService{repo: repo}
}

type IUserRepo interface {
	Create(ctx context.Context, user *User) (int, error)
}

func (us *UserService) Create(ctx context.Context, dto *CreateUserDTO) (int, error) {
	user := NewUserFromCreateDTO(dto)
	if err := user.generatePasswordHash(); err != nil {
	}
	userID, err := us.repo.Create(ctx, user)
	if err != nil {
	}
}
