package main

import (
	"jwt-auth/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// turn off debug mode
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)

	r.Run(":8080")
}
