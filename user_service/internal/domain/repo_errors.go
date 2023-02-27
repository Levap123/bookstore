package domain

import "errors"

var (
	ErrUnique       = errors.New("user with this username or email already exist")
	ErrUserNotFound = errors.New("user with this email not found")
)
