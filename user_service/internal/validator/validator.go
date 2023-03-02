package validator

import (
	"net/mail"
	"regexp"

	"github.com/Levap123/user_service/internal/configs"
)

type Validator struct {
	PasswordMax int
	PasswordMin int
	emailReg    *regexp.Regexp

	UsernameMin int
	UsernameMax int
}

func NewValidator(cfg *configs.Configs) *Validator {
	return &Validator{
		PasswordMin: cfg.Validator.PasswordMin,
		PasswordMax: cfg.Validator.PasswordMax,
		emailReg:    regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`),
		UsernameMin: cfg.Validator.UsernameMin,
		UsernameMax: cfg.Validator.UsernameMax,
	}
}

func (v *Validator) IsUsernameLengthCorrect(Username string) bool {
	return len([]rune(Username)) >= v.UsernameMin && len([]rune(Username)) <= v.UsernameMax
}

func (v *Validator) IsPasswordLenghtCorrect(password string) bool {
	return len([]rune(password)) >= v.PasswordMin && len([]rune(password)) <= v.PasswordMax
}

func (v *Validator) IsEmailCorrect(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
