package api

import (
	"github.com/gin-gonic/gin"
	"github.com/karma-dev-team/karma-docs/internal/user"
)

func RegisterUser(router *gin.Engine, userService user.UserServcie) {
	userApi := router.Group("/user")

	userRouter := UserRouter{
		userService: userService,
	}

	userApi.GET("/user/:userid", userRouter.GetUserById)
	userApi.DELETE("/user/:userid", userRouter.DeleteUser)
}
