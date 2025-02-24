package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "devops-challenge/docs"
	"devops-challenge/handlers"
)

// @title DevOps Challenge API
// @version 1.0
// @description This is a simple API for the DevOps challenge.
// @host localhost:8080
// @BasePath /


func main() {
	// Load environment variables
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

	// Setup Swagger handler
	handlers.SetupSwagger(router)

	// Health check route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Define API routes
	router.GET("/deals", handlers.GetDeals)
	// router.POST("/deals", handlers.CreateDeal)
	// router.PUT("/deals", handlers.UpdateDeal)


	// Start server on port 8080
	log.Println("Server is running on port 8080...")
	router.Run(":8080")
}