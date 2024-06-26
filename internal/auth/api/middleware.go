package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/karma-dev-team/karma-docs/internal/auth"
)

type AuthMiddleware struct {
	authService auth.AuthService
}

func NewAuthMiddleware(service auth.AuthService) gin.HandlerFunc {
	return (&AuthMiddleware{
		authService: service,
	}).Handle
}

func (m *AuthMiddleware) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := m.authService.ParseToken(c.Request.Context(), headerParts[1])
	if err != nil {
		status := http.StatusInternalServerError
		if err == auth.ErrInvalidAccessToken {
			status = http.StatusUnauthorized
		}

		c.AbortWithStatus(status)
		return
	}

	c.Set(auth.CtxUserKey, user)
}
