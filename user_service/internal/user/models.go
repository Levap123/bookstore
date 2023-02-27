package user

import (
	"user_service/proto"

	"github.com/Levap123/utils/crypt"
)

type User struct {
	ID       uint64
	Email    string
	Username string
	Password string
}

type CreateUserDTO struct {
	Email    string
	Username string
	Password string
}

type GetUserDTO struct {
	Email    string
	Password string
}

type UpdateUserDTO struct {
	ID       int
	Username string
	Password string
}

func NewUpdateUserDTO(pb *proto.UpdateUserRequest) *UpdateUserDTO {
	return &UpdateUserDTO{
		Username: pb.Username,
		Password: pb.Password,
	}
}

func NewGetUserDTO(pb *proto.SignInRequest) *GetUserDTO {
	return &GetUserDTO{
		Email:    pb.Email,
		Password: pb.Password,
	}
}

func NewCreateUserDTO(pb *proto.SignUpRequest) *CreateUserDTO {
	return &CreateUserDTO{
		Email:    pb.Email,
		Username: pb.Username,
		Password: pb.Password,
	}
}

func NewUserFromCreateDTO(dto *CreateUserDTO) *User {
	return &User{
		Email:    dto.Email,
		Username: dto.Username,
		Password: dto.Password,
	}
}

func NewUserFromUpdateDTO(dto *UpdateUserDTO) *User {
	return &User{
		ID:       uint64(dto.ID),
		Username: dto.Username,
		Password: dto.Password,
	}
}

func (user *User) generatePasswordHash() error {
	pwd, err := crypt.GeneratePasswordHash(user.Password)
	if err != nil {
		return err
	}
	user.Password = pwd
	return nil
}

func (user *User) PasswordCorrect(password string) bool {
	return crypt.ComparePassword(password, user.Password) == nil
}
