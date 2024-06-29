package server

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/karma-dev-team/karma-docs/internal/auth"
	"github.com/karma-dev-team/karma-docs/internal/auth/api"
	authApi "github.com/karma-dev-team/karma-docs/internal/auth/api"
	"github.com/karma-dev-team/karma-docs/internal/auth/usecases"
	"github.com/karma-dev-team/karma-docs/internal/config"
	"github.com/karma-dev-team/karma-docs/internal/docs"
	docsApi "github.com/karma-dev-team/karma-docs/internal/docs/api"
	docsRepositories "github.com/karma-dev-team/karma-docs/internal/docs/repositories"
	docsUsecase "github.com/karma-dev-team/karma-docs/internal/docs/usecases"
	"github.com/karma-dev-team/karma-docs/internal/user"
	userApi "github.com/karma-dev-team/karma-docs/internal/user/api"
	"github.com/karma-dev-team/karma-docs/internal/user/repositories"
	userUsecase "github.com/karma-dev-team/karma-docs/internal/user/usecases"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	httpServer *http.Server

	authService auth.AuthService
	docsService docs.DocumentService
	userService user.UserServcie

	Logger *slog.Logger
}

func NewApp(config *config.AppConfig) *App {
	loggerInstance := initLogging(config)
	db := initDB(config, loggerInstance)

	userRepo := repositories.NewUserRepository(db)
	docsRepo := docsRepositories.NewDocumentRepository(db)

	return &App{
		authService: usecases.NewAuthService(
			userRepo,
			[]byte(config.Jwt.TokenKey),
			time.Duration(config.Jwt.ExpireDuration),
		),
		docsService: docsUsecase.NewDocumentService(docsRepo),
		userService: userUsecase.NewUserService(userRepo),
		Logger:      loggerInstance,
	}
}

func (a *App) Run(port string) error {
	// Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	authMiddleware := api.NewAuthMiddleware(a.authService)
	api := router.Group("/api", authMiddleware)

	// Set up http handlers
	// SignUp/SignIn endpoints
	authApi.RegisterAuth(router, a.authService)
	userApi.RegisterUser(router, a.userService)
	docsApi.RegisterDocs(api, a.docsService)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func initDB(config *config.AppConfig, logger *slog.Logger) *gorm.DB {
	dsn := config.GenerateDSN()
	dbConfig := &gorm.Config{}
	logger.Info("Creating database, opening database")
	db, err := gorm.Open(postgres.Open(dsn), dbConfig)
	if err != nil {
		// early panic, so it cannot continue
		panic(err)
	}
	return db
}

func initLogging(config *config.AppConfig) *slog.Logger {
	logDir := filepath.Dir(config.Logger.Path)
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		// crash
		panic(errors.Wrapf(err, "failed to create log directory: %s", logDir))
	}

	// Open the log file
	file, err := os.OpenFile(config.Logger.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a JSON handler for the logger
	jsonHandler := slog.NewJSONHandler(file, &slog.HandlerOptions{})

	// Create a new logger with the JSON handler
	logger := slog.New(jsonHandler)

	slog.SetDefault(logger)

	return logger
}
