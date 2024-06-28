package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karma-dev-team/karma-docs/internal/auth"
)

type AuthHandler struct {
	authService auth.AuthService
}

type signInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type signUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type signInResponse struct {
	Token string `json:"token"`
}

func (a *AuthHandler) SignIn(c *gin.Context) {
	var signinRequest signInRequest
	if err := c.BindJSON(&signinRequest); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := a.authService.SignIn(c.Request.Context(), signinRequest.Username, signinRequest.Password)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, signInResponse{Token: token})
}

func (a *AuthHandler) SingUp(c *gin.Context) {
	inp := new(signUpRequest)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := a.authService.SignUp(c.Request.Context(), inp.Email, inp.Username, inp.Password); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
