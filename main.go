package main

import (
	"jwt-auth/controllers"
	"jwt-auth/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// turn off debug mode
	gin.SetMode(gin.ReleaseMode)

	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)

	r.Run(":8080")
}
