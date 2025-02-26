package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"

	_ "devops-challenge/docs"
	"devops-challenge/handlers"
)

// Logger instance with JSON formatting
var logger = logrus.New()

// Custom Prometheus registry (excluding default Go metrics)
var customRegistry = prometheus.NewRegistry()

// Prometheus metrics for API requests
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
			Name:    "api_request_duration_milliseconds",
			Help:    "Histogram of API response time in milliseconds",
			Buckets: prometheus.LinearBuckets(5, 10, 10), // Buckets from 5ms, increasing by 10ms
		},
		[]string{"method", "path"},
	)
)

func init() {
	// Use JSON formatter for clean logging
	logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true, // Makes logs easier to read
	})

	// Register only API-specific metrics
	customRegistry.MustRegister(requestCount, requestDuration)
}

// Middleware to track API metrics with structured logging
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Generate a unique request ID for tracking
		requestID := uuid.New().String()
		c.Set("requestID", requestID)

		// Process the request
		c.Next()

		// Compute response time in milliseconds
		duration := time.Since(start).Milliseconds()
		statusCode := c.Writer.Status()

		// Log structured API request details
		logger.WithFields(logrus.Fields{
			"request_id":   requestID,
			"method":       c.Request.Method,
			"path":         c.Request.URL.Path,
			"status_code":  statusCode,
			"latency_ms":   duration,
			"client_ip":    c.ClientIP(),
			"user_agent":   c.Request.UserAgent(),
			"content_type": c.Request.Header.Get("Content-Type"),
		}).Info("API request completed")

		// Update Prometheus metrics
		requestCount.WithLabelValues(c.Request.Method, c.Request.URL.Path, fmt.Sprintf("%d", statusCode)).Inc()
		requestDuration.WithLabelValues(c.Request.Method, c.Request.URL.Path).Observe(float64(duration))
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
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Define API routes
	apiV1 := router.Group("/v1")
	{
		apiV1.GET("/deals", handlers.GetAllDeals)
		apiV1.POST("/deals", handlers.CreateDeal)
		apiV1.PUT("/deals/:id", handlers.UpdateDeal)
	}

	// Expose ONLY API-specific metrics
	router.GET("/metrics", gin.WrapH(promhttp.HandlerFor(customRegistry, promhttp.HandlerOpts{})))

	logger.Info("Server is running on port 8080...")
	router.Run(":8080")
}
