package api

import (
	"github.com/gin-gonic/gin"
	"github.com/karma-dev-team/karma-docs/internal/auth"
)

func RegisterAuth(router *gin.Engine, authService auth.AuthService) {
	router.Use(NewAuthMiddleware(authService))

	authApi := router.Group("/auth")
	authHandler := AuthHandler{
		authService: authService,
	}

	authApi.POST("/login", authHandler.SignIn)
}
