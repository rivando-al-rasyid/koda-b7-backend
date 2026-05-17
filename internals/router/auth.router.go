package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rivando-al-rasyid/koda-b7-backend/internals/controller"
)

func RegisterAuthRoutes(r *gin.Engine) {
	ac := controller.NewAuthController()

	r.POST("/login", ac.Login)
	r.POST("/register", ac.Register)
}
