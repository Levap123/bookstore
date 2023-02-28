package domain

import "errors"

var (
	ErrUnique       = errors.New("user with this username or email already exists")
	ErrUserNotFound = errors.New("user not found")
)
