package domain

import "errors"

var (
	ErrPasswordLengthIncorrect = errors.New("password length incorrect")
	ErrUsernameLengthIncorrect = errors.New("username length incorrect")
	ErrIncorrectEmail          = errors.New("incorrect email format")
)
