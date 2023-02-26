package user

import (
	"context"
	"fmt"
)

type UserService struct {
	repo IUserRepo
}

func NewUserService(repo IUserRepo) *UserService {
	return &UserService{repo: repo}
}

type IUserRepo interface {
	Create(ctx context.Context, user *User) (uint64, error)
}

func (us *UserService) Create(ctx context.Context, dto *CreateUserDTO) (uint64, error) {
	user := NewUserFromCreateDTO(dto)
	if err := user.generatePasswordHash(); err != nil {
	}
	userID, err := us.repo.Create(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("user service - repo create - %w", err)
	}
	return userID, nil
}
