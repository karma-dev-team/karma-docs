package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/karma-dev-team/karma-docs/internal/user"
)

type UserRouter struct {
	userService *user.UserServcie
}

type getUserById struct {
	Id uuid.UUID `json:"id"`
}

func (user *UserRouter) GetUserById(c *gin.Context) {
	req := new(getUserById)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.MustGet()
}
