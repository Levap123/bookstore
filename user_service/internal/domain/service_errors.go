package domain

import "errors"

var (
	ErrIncorrectPassword  = errors.New("user password incorrect")
	ErrIncorrectTokenType = errors.New("token type incorrect")
	ErrTokensMissmatched  = errors.New("this token belongs to different users")
)
