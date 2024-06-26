package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/karma-dev-team/karma-docs/internal/user/entities"
)

type GetUserRequest struct {
	Username string
	Email    string
	UserId   uuid.UUID
}

type UserRepository interface {
	AddUser(ctx context.Context, user *entities.User) error
	GetUser(ctx context.Context, request GetUserRequest) (*entities.User, error)
	DeleteUser(ctx context.Context, userId uuid.UUID) error
}
