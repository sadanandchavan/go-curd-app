package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sadanandchavan/go-crud-app/config"
	"github.com/sadanandchavan/go-crud-app/handlers"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	// ✅ CORS middleware must be added before routes
	r.Use(cors.Default())

	r.POST("/properties", handlers.CreateProperty)
	r.GET("/properties", handlers.GetProperties)
	r.GET("/properties/:id", handlers.GetProperty)
	r.PUT("/properties/:id", handlers.UpdateProperty)
	r.DELETE("/properties/:id", handlers.DeleteProperty)

	r.Run(":8080") // localhost:8080
}
