package api

import "github.com/gin-gonic/gin"

func RegisterUser(router *gin.Engine) {
	user_api := router.Group("/user")

	user_api.Get()
}
