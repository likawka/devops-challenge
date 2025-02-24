package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"devops-challenge/handlers"
)

func main() {
	// Load .env file (ignores errors if the file is missing)
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	// Read API key
	apiKey, exists := os.LookupEnv("PIPEDRIVE_API_KEY")
	if !exists || apiKey == "" {
		log.Fatal("PIPEDRIVE_API_KEY is not set")
	}

	// Initialize Gin router
	router := gin.Default()

	// Health check route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Define routes
	router.GET("/deals", handlers.GetDeals)
	router.POST("/deals", handlers.CreateDeal)
	router.PUT("/deals", handlers.UpdateDeal)

	// Start server on port 8080
	log.Println("Server is running on port 8080...")
	router.Run(":8080")
}
