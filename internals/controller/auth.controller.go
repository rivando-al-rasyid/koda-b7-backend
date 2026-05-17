package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rivando-al-rasyid/koda-b7-backend/internals/dto"
	"github.com/rivando-al-rasyid/koda-b7-backend/internals/service"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		authService: service.NewAuthService(),
	}
}

func (ac *AuthController) Login(c *gin.Context) {
	var logreq dto.LoginRequest
	if err := c.ShouldBindJSON(&logreq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !ac.authService.Checkmail(logreq.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email format",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"email":   logreq.Email,
	})
}

func (ac *AuthController) Register(c *gin.Context) {
	var regisreq dto.RegisterRequest
	if err := c.ShouldBindJSON(&regisreq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !ac.authService.Checkmail(regisreq.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email format",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "register success",
		"email":   regisreq.Email,
	})
}
