package main

import (
	"jwt-auth/controllers"
	"jwt-auth/middlewares"
	"jwt-auth/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// turn off debug mode
	gin.SetMode(gin.ReleaseMode)

	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	log.Println("Server started on: ", "http://127.0.0.1:8080")
	r.Run(":8080")
}
