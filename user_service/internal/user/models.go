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

func (user *User) generatePasswordHash() error {
	pwd, err := crypt.GeneratePasswordHash(user.Password)
	if err != nil {
		return err
	}
	user.Password = pwd
	return nil
}
