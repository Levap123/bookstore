package postgres_test

import (
	"context"
	"io"
	"reflect"

	"testing"
	"user_service/internal/user"
	"user_service/internal/user/postgres"

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
	}
)

func TestCreate(t *testing.T) {
	repo := postgres.NewUserRepo(DB, log)

	type args struct {
		ctx  context.Context
		user *user.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    uint64
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
			name: "should create user with a error",
			args: args{
				ctx:  context.Background(),
				user: users[0],
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.Create(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.Create(), err = %v, wantErr %v", err, tt.wantErr)
				return
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
		name    string
		args    args
		wantErr bool
		want    *user.User
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
			name: "should throw error",
			args: args{
				ctx:   context.Background(),
				email: "does not exist@gmail.com",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetByEmail(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetByEmail(), err = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if reflect.DeepEqual(got, users[0]) {
				t.Errorf("UserRepository.GetByEmail(), expected = %v, got %v", tt.want, got)
			}
		})
	}
}
