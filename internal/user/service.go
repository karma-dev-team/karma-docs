package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/karma-dev-team/karma-docs/internal/user/entities"
)

type UserServcie interface {
	CreateUser(ctx context.Context, s string) (*entities.User, error)
	GetUser(ctx context.Context, userId uuid.UUID) (*entities.User, error)
}
