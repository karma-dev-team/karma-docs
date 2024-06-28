package server

import (
	"net/http"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/karma-dev-team/karma-docs/internal/auth"
	"github.com/karma-dev-team/karma-docs/internal/auth/usecases"
	"github.com/karma-dev-team/karma-docs/internal/config"
	"github.com/karma-dev-team/karma-docs/internal/docs"
	"github.com/karma-dev-team/karma-docs/internal/user"
	"github.com/karma-dev-team/karma-docs/internal/user/repositories"
)

type App struct {
	httpServer *http.Server

	authService auth.AuthService
	docsService docs.DocumentService
	userService user.UserServcie
}

func NewApp(pool *pgxpool.Pool, config *config.AppConfig) *App {
	userRepo := repositories.NewUserRepository(pool)

	return &App{
		authService: usecases.NewAuthService(
			userRepo,
			[]byte(config.Jwt.TokenKey),
			time.Duration(config.Jwt.ExpireDuration),
		),
		docsService: usecases.NewDocsService(),
	}
}
