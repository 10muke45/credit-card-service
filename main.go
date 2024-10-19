package main

import (
	"kreval/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Route for card validation
	router.POST("/validate-card", handlers.ValidateCard)

	// Route for card number generation
	router.POST("/generate-card", handlers.GenerateCard)

	router.Run(":8080") // Start server on port 8080
}
