package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rivando-al-rasyid/koda-b7-backend/internals/router"
)

func main() {
	app := gin.Default()

	router.RegisterAuthRoutes(app)

	app.Run("localhost:8080")
}
