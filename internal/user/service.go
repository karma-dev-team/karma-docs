package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/karma-dev-team/karma-docs/internal/user/entities"
)

type CreateUserDto struct {
	Username string
	Password string
	Email    string
}

type UserServcie interface {
	CreateUser(ctx context.Context, dto CreateUserDto) error
	GetUser(ctx context.Context, userId uuid.UUID) (*entities.User, error)
	DeleteUser(ctx context.Context, userId uuid.UUID) error
}
