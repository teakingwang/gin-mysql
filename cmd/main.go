package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"my-gin-app/config"
	"my-gin-app/internal/app"
)

func main() {
	cfg := config.LoadConfig()

	r := gin.Default()

	// Register routes
	userRoutes := r.Group("/api/v1/users")
	{
		userRoutes.GET("", app.GetUserList)
		userRoutes.POST("", app.CreateUser)
	}

	// Start server
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
