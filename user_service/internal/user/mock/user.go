package mock

import (
	"context"

	"user_service/internal/domain"
	"user_service/internal/user"
)

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

const userTable = "users"

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

func Create(ctx context.Context, user *user.User) (uint64, error) {
	for _, userIn := range users {
		if userIn.Email == user.Email || userIn.Username == userIn.Username {
			return 0, domain.ErrUnique
		}
	}
	user.ID = uint64(len(users) + 1)
	users = append(users, user)
	return user.ID, nil
}

func GetByEmail(ctx context.Context, email string) (*user.User, error) {
	for _, userIn := range users {
		if userIn.Email == email {
			return userIn, nil
		}
	}
	return nil, domain.ErrUserNotFound
}

func GetByID(ctx context.Context, ID uint64) (*user.User, error) {
	for _, userIn := range users {
		if userIn.ID == ID {
			return userIn, nil
		}
	}
	return nil, domain.ErrUserNotFound
}

func UpdateInfo(ctx context.Context, user *user.User) (int, error) {
	for ind, userIn := range users {
		if userIn.ID == user.ID {
			users[ind] = user
			return ind + 1, nil
		}
	}
	return 0, domain.ErrUserNotFound
}
