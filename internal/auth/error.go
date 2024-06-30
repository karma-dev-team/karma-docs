package auth

import "errors"

var (
	ErrInvalidAccessToken = errors.New("Invalid access token")
	ErrUserNotFound       = errors.New("User not found")
	ErrAccessDenied       = errors.New("Access denied")
)
