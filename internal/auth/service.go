package auth

import (
	"context"

	"github.com/karma-dev-team/karma-docs/internal/user/entities"
)

const CtxUserKey = "user"

type AuthService interface {
	SignUp(ctx context.Context, username, password string) error
	SignIn(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*entities.User, error)
}
