package postgres_test

import (
	"context"
	"errors"
	"io"
	"reflect"
	"testing"

	"github.com/Levap123/user_service/internal/domain"
	"github.com/Levap123/user_service/internal/user"
	"github.com/Levap123/user_service/internal/user/postgres"

	"github.com/sirupsen/logrus"
)

var (
	log = &logrus.Logger{
		Out: io.Discard,
	}

	users = []*user.User{
		{
			Email:    "test@mail.ru",
			Username: "test",
			Password: "testpass",
		},
		{
			Email:    "unique@gmail.com",
			Username: "unique",
			Password: "uniquepass",
		},
		{
			Email:    "john@gmail.com",
			Username: "john",
			Password: "jonhpass",
		},
	}
)

func TestCreate(t *testing.T) {
	repo := postgres.NewUserRepo(DB, log)

	type args struct {
		ctx  context.Context
		user *user.User
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		wantThisErr error
		want        uint64
	}{
		{
			name: "should create user without any error",
			args: args{
				ctx:  context.Background(),
				user: users[0],
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "should create user without any error",
			args: args{
				ctx:  context.Background(),
				user: users[1],
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "should create user without any error",
			args: args{
				ctx:  context.Background(),
				user: users[2],
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "should create user with a error",
			args: args{
				ctx:  context.Background(),
				user: users[0],
			},
			wantThisErr: domain.ErrUnique,
			wantErr:     true,
		},
		{
			name: "should create user with a error",
			args: args{
				ctx:  context.Background(),
				user: users[1],
			},
			wantThisErr: domain.ErrUnique,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.Create(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.Create(), err = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if !errors.Is(err, tt.wantThisErr) {
					t.Errorf("UserRepository.Create(), expected err = %v, got %v", tt.wantThisErr, err)
					return
				}
			}
			if got != tt.want {
				t.Errorf("UserRepository.Create(), expected = %v, got %v", tt.want, got)
				return
			}
		})
	}
}

func TestGetByEmail(t *testing.T) {
	repo := postgres.NewUserRepo(DB, log)

	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		wantThisErr error
		want        *user.User
	}{
		{
			name: "should return user without error",
			args: args{
				ctx:   context.Background(),
				email: "test@mail.ru",
			},
			want:    users[0],
			wantErr: false,
		},
		{
			name: "should throw error 'user not found'",
			args: args{
				ctx:   context.Background(),
				email: "does not exist@gmail.com",
			},
			wantErr:     true,
			wantThisErr: domain.ErrUserNotFound,
		},
		{
			name: "should throw error 'user not found'",
			args: args{
				ctx:   context.Background(),
				email: "123123@gmail.com",
			},
			wantErr:     true,
			wantThisErr: domain.ErrUserNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetByEmail(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetByEmail(), err = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if !errors.Is(err, tt.wantThisErr) {
					t.Errorf("UserRepository.GetByEmail(), expected err = %v, got %v", tt.wantThisErr, err)
					return
				}
			}
			if reflect.DeepEqual(got, users[0]) {
				t.Errorf("UserRepository.GetByEmail(), expected = %v, got %v", tt.want, got)
			}
		})
	}
}

func TestGetByID(t *testing.T) {
	repo := postgres.NewUserRepo(DB, log)

	type args struct {
		ctx    context.Context
		userID uint64
	}

	tests := []struct {
		name        string
		args        args
		wantErr     bool
		wantThisErr error
		want        *user.User
	}{
		{
			name: "should get user without any error",
			args: args{
				context.Background(),
				1,
			},
			wantErr: false,
			want:    users[0],
		},
		{
			name: "should get user without any error",
			args: args{
				context.Background(),
				2,
			},
			wantErr: false,
			want:    users[2],
		},
		{
			name: "should throw error 'user not found'",
			args: args{
				context.Background(),
				100,
			},
			wantErr:     true,
			wantThisErr: domain.ErrUserNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetByID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetByID(), expected = %v, wantErr %v", tt.want, tt.wantErr)
				return
			}

			if tt.wantErr {
				if !errors.Is(err, domain.ErrUserNotFound) {
					t.Errorf("UserRepository.GetByID(), expected err = %v, got %v", tt.want, got)
					return
				}
			}

			if reflect.DeepEqual(got, users[0]) {
				t.Errorf("UserRepository.GetByID(), expected = %v, got %v", tt.want, got)
			}
		})
	}
}

func TestUpdateInfo(t *testing.T) {
	repo := postgres.NewUserRepo(DB, log)

	type args struct {
		ctx  context.Context
		user *user.User
	}

	tests := []struct {
		name        string
		args        args
		wantErr     bool
		wantThisErr error
		want        int
	}{
		{
			name: "should update user with error 'unique'",
			args: args{
				ctx: context.Background(),
				user: &user.User{
					Username: "unique",
					Email:    "john@gmail.com",
					Password: "new password",
				},
			},
			wantErr:     true,
			wantThisErr: domain.ErrUnique,
		},
		{
			name: "should update user without error",
			args: args{
				ctx: context.Background(),
				user: &user.User{
					ID:       1,
					Username: "new username",
					Email:    "unique@gmail.com",
					Password: "new password",
				},
			},
			want:    2,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.UpdateInfo(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.UpdateInfo(), expected = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				if !errors.Is(err, tt.wantThisErr) {
					t.Errorf("UserRepository.UpdateInfo(), expected err = %v, got  %v", tt.wantThisErr, err)
					return
				}
			}

			if got != tt.want {
				t.Errorf("UserRepository.UpdateInfo(), expected  = %v, got  %v", tt.want, got)
				return
			}
		})
	}
}
