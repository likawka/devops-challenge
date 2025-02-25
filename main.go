package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"

	_ "devops-challenge/docs"
	"devops-challenge/handlers"
)

// Logger instance
var logger = logrus.New()

// Create a new Prometheus registry (to disable default Go metrics)
var customRegistry = prometheus.NewRegistry()

// API-specific Prometheus metrics
var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_requests_total",
			Help: "Total number of API requests",
		},
		[]string{"method", "path", "status"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_request_duration_seconds",
			Help:    "Histogram of API response time",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)
)

func init() {
	// Set log format to JSON for better readability
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Register only API metrics (skip default Go metrics)
	customRegistry.MustRegister(requestCount, requestDuration)
}

// Middleware to track API metrics
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		duration := time.Since(start).Seconds()
		statusCode := fmt.Sprintf("%d", c.Writer.Status())

		// Log API request details
		logger.WithFields(logrus.Fields{
			"method":      c.Request.Method,
			"path":        c.Request.URL.Path,
			"status":      statusCode,
			"latency_sec": duration,
			"client_ip":   c.ClientIP(),
		}).Info("API request completed")

		// Update Prometheus API metrics
		requestCount.WithLabelValues(c.Request.Method, c.Request.URL.Path, statusCode).Inc()
		requestDuration.WithLabelValues(c.Request.Method, c.Request.URL.Path).Observe(duration)
	}
}

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		logger.Warn("No .env file found, using system environment variables")
	}

	// Read API key
	apiKey, exists := os.LookupEnv("PIPEDRIVE_API_KEY")
	if !exists || apiKey == "" {
		logger.Fatal("PIPEDRIVE_API_KEY is not set")
	}

	// Initialize Gin router
	router := gin.Default()

	// Apply metrics middleware
	router.Use(MetricsMiddleware())

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
		apiV1.POST("/deals", handlers.CreateDeal)
		apiV1.PUT("/deals/:id", handlers.UpdateDeal)
	}

	router.GET("/metrics", gin.WrapH(promhttp.HandlerFor(customRegistry, promhttp.HandlerOpts{})))

	logger.Info("Server is running on port 8080...")
	router.Run(":8080")
}
