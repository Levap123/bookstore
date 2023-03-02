package mock

import (
	"context"

	"github.com/Levap123/user_service/internal/domain"
	"github.com/Levap123/user_service/internal/user"
)

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

var users = []*user.User{
	{
		ID:       1,
		Email:    "arturpidor@mail.ru",
		Username: "arturpidor",
		Password: "arturpidor",
	},
	{
		ID:       2,
		Email:    "test@mail.ru",
		Username: "test",
		Password: "test",
	},
	{
		ID:       3,
		Email:    "levap@gmail.com",
		Username: "levap",
		Password: "levap",
	},
}

func (ur *UserRepo) Create(ctx context.Context, user *user.User) (uint64, error) {
	for _, userIn := range users {
		if userIn.Email == user.Email || userIn.Username == user.Username {
			return 0, domain.ErrUnique
		}
	}
	user.ID = uint64(len(users) + 1)
	users = append(users, user)
	return user.ID, nil
}

func (ur *UserRepo) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	for _, userIn := range users {
		if userIn.Email == email {
			return userIn, nil
		}
	}
	return nil, domain.ErrUserNotFound
}

func (ur *UserRepo) GetByID(ctx context.Context, ID uint64) (*user.User, error) {
	for _, userIn := range users {
		if userIn.ID == ID {
			return userIn, nil
		}
	}
	return nil, domain.ErrUserNotFound
}

func (ur *UserRepo) UpdateInfo(ctx context.Context, user *user.User) (int, error) {
	for ind, userIn := range users {
		if userIn.Email == user.Email {
			user.Email = users[ind].Email
			users[ind] = user
			return ind + 1, nil
		}
	}
	return 0, domain.ErrUserNotFound
}
