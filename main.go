package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "devops-challenge/docs"
	"devops-challenge/handlers"
)

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
	apiV1 := router.Group("/v1")
	{
		apiV1.GET("/deals", handlers.GetAllDeals)

	}

	log.Println("Server is running on port 8080...")
	router.Run(":8080")
}
