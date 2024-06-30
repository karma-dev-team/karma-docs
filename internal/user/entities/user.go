package entities

import (
	"github.com/karma-dev-team/karma-docs/internal/security"
	"github.com/karma-dev-team/karma-docs/pkg/gormplugin"
)

type User struct {
	gormplugin.Model
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	IsBlocked      bool   `json:"is_blocked"`
}

type UserDomainService struct {
}

func NewUser(username string, email string, password string) (*User, error) {
	hashed_password, err := security.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := User{
		Username:       username,
		Email:          email,
		HashedPassword: hashed_password,
		IsBlocked:      false,
	}

	return &user, nil
}
