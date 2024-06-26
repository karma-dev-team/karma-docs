package entities

import (
	"github.com/google/uuid"
	"github.com/karma-dev-team/karma-docs/internal/security"
)

type User struct {
	Id             uuid.UUID `json: "id"`
	Username       string    `json: "username"`
	Email          string    `json: "email"`
	HashedPassword string    `json: "hashed_password"`
}

type UserDomainService struct {
}

func NewUser(username string, email string, password string) (*User, error) {
	hashed_password, err := security.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := User{
		Id:             uuid.New(),
		Username:       username,
		Email:          email,
		HashedPassword: hashed_password,
	}

	return &user, nil
}
