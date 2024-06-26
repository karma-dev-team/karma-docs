package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/karma-dev-team/karma-docs/internal/user"
)

type UserRouter struct {
	userService user.UserServcie
}

func (u *UserRouter) GetUserById(c *gin.Context) {

	userId, err := uuid.Parse(c.Param("userid"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, user.ErrUserNotFound)
		return
	}
	user, err := u.userService.GetUser(c.Request.Context(), userId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *UserRouter) DeleteUser(c *gin.Context) {
	userId, err := uuid.Parse(c.Param("userid"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = u.userService.DeleteUser(c.Request.Context(), userId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Status(http.StatusOK)
}
