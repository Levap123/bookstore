package validator

import (
	"regexp"
	"user_service/internal/configs"
)

type Validator struct {
	passwordMin int
	passwordMax int

	emailReg *regexp.Regexp

	usernameMin int
	usernameMax int
}

func NewValidator(cfg *configs.Configs) *Validator {
	return &Validator{
		passwordMin: cfg.Validator.PasswordMin,
		passwordMax: cfg.Validator.PasswordMax,
		emailReg:    regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`),
		usernameMin: cfg.Validator.UsernameMin,
		usernameMax: cfg.Validator.UsernameMax,
	}
}
