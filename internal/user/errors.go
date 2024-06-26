package user

import "errors"

var (
	ErrUserNotFound = errors.New("User with ginven id not found")
)
