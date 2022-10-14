package routes

import (
	"github.com/gin-gonic/gin"
	"task/api/handlers"
)

func SetupRoutes() {
	// Root route
	root := gin.Default()

	// Route group
	client := root.Group("/client")

	// Check balances route
	client.POST("/start", handlers.StartTransaction)
	client.POST("/add", handlers.AddBalance)

	// Rout route
	root.Run()
}
