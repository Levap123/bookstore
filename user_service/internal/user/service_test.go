package user_test

import (
	"context"
	"testing"

	"user_service/internal/user"
	"user_service/internal/user/mock"

	"github.com/Levap123/utils/jwt"
)

var us = user.NewUserService(mock.NewUserRepo(), jwt.NewJWT("testingSign"))

var userCreateDTOs = []*user.CreateUserDTO{
	{
		Email:    "unique",
		Username: "unique",
		Password: "password",
	},
	{
		Email:    "unique",
		Username: "arturpidor",
		Password: "password",
	},
}

func TestUserService_Create(t *testing.T) {
	type args struct {
		ctx context.Context
		dto *user.CreateUserDTO
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "should create user without any error",
			args: args{
				context.Background(),
				userCreateDTOs[0],
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "should create user with error",
			args: args{
				context.Background(),
				userCreateDTOs[1],
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := us.Create(tt.args.ctx, tt.args.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UserService.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

var getUserDTOs = []*user.GetUserDTO{
	{
		Email:    "unique",
		Password: "password",
	},
	{
		Email:    "unique",
		Password: "passwordincorrect",
	},
}

func TestUserService_GenerateTokens(t *testing.T) {
	type args struct {
		ctx context.Context
		dto *user.GetUserDTO
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		{
			name: "should sign in without any error",
			args: args{
				context.Background(),
				getUserDTOs[0],
			},
			wantErr: false,
		},
		{
			name: "should signin with error incorrect password",
			args: args{
				context.Background(),
				getUserDTOs[1],
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := us.GenerateTokens(tt.args.ctx, tt.args.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.GenerateTokens() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				_, err = us.Validate(context.Background(), got)
				if err != nil {
					t.Errorf("UserService.GenerateTokens() error = %v, want nil", err)
					return
				}
			}
		})
	}
}

var updateUserDTOs = []*user.UpdateUserDTO{
	{
		ID:          4,
		Username:    "unique",
		OldPassword: "password",
		NewPassword: "testest",
	},
	{
		ID:          4,
		Username:    "unique",
		OldPassword: "p12assword",
		NewPassword: "testest",
	},
}

func TestUserService_UpdateUser(t *testing.T) {
	type args struct {
		ctx context.Context
		dto *user.UpdateUserDTO
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "should update user without any error",
			args: args{
				context.Background(),
				updateUserDTOs[0],
			},
			wantErr: false,
			want:    4,
		},
		{
			name: "should update user wit error",
			args: args{
				context.Background(),
				updateUserDTOs[1],
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := us.UpdateUser(tt.args.ctx, tt.args.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UserService.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
