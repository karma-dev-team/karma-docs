package user

import "errors"

var (
	ErrUserNotFound  = errors.New("User with ginven id not found")
	ErrUserIsBlocked = errors.New("User with given id is blocked")
)
