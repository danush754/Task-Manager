package handler

import (
	"github.com/danush754/Task-Manager/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	UserService *service.UserService
}

func NewAuthHandler(service *service.UserService) *AuthHandler {
	return &AuthHandler{UserService: service}
}

func (a *AuthHandler) SignUp(c *gin.Context) {

}

func (a *AuthHandler) Login(c *gin.Context) {

}
