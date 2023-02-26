package domain

import "errors"

var ErrIncorrectPassword = errors.New("user password incorrect")
var ErrIncorrectTokenType = errors.New("token type must be access")