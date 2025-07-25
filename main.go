package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sadanandchavan/go-crud-app/config"
	"github.com/sadanandchavan/go-crud-app/handlers"
)

func main() {
	config.ConnectDB()
	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://go-crud-frontend-ten.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Handle OPTIONS
	r.OPTIONS("/*path", func(c *gin.Context) {
		c.AbortWithStatus(204)
	})

	// API routes
	r.POST("/properties", handlers.CreateProperty)
	r.GET("/properties", handlers.GetProperties)
	r.GET("/properties/:id", handlers.GetProperty)
	r.PUT("/properties/:id", handlers.UpdateProperty)
	r.DELETE("/properties/:id", handlers.DeleteProperty)

	r.Run(":8080")
}
