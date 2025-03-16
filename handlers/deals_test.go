package handlers

import (
	"bytes"
	"devops-challenge/api"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var apiKey string

// Load environment variables before running tests
func TestMain(m *testing.M) {
	_ = godotenv.Load(".env") // Load .env file
	apiKey = os.Getenv("PIPEDRIVE_API_KEY") // Set API key
	os.Exit(m.Run()) // Run tests
}

// setupRouter initializes the Gin router with endpoints for testing
func setupRouter() *gin.Engine {
	router := gin.Default()

	// Mock handlers instead of making real API calls
	router.POST("/v1/deals", mockCreateDeal)
	router.PUT("/v1/deals/:id", mockUpdateDeal)
	router.GET("/v1/deals/:id", mockGetDeal)

	return router
}

// Mock CreateDeal handler
func mockCreateDeal(c *gin.Context) {
	var request map[string]interface{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid JSON"})
		return
	}

	// Simulated success response
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data": map[string]interface{}{
			"id":       1,
			"title":    request["title"],
			"value":    request["value"],
			"currency": request["currency"],
		},
	})
}

// Mock UpdateDeal handler
func mockUpdateDeal(c *gin.Context) {
	var request map[string]interface{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid JSON"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": map[string]interface{}{
			"id":    c.Param("id"),
			"title": request["title"],
		},
	})
}

// Mock GetDeal handler
func mockGetDeal(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": map[string]interface{}{
			"id":       c.Param("id"),
			"title":    "Existing Deal",
			"value":    100,
			"currency": "USD",
		},
	})
}

// TestCreateDeal checks if a new deal can be created
func TestCreateDeal(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	dealJSON := `{"title": "New Deal", "value": 100, "currency": "USD"}`
	req, _ := http.NewRequest("POST", "/v1/deals", bytes.NewBufferString(dealJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, true, response["success"])
	assert.NotNil(t, response["data"])
}

// TestUpdateDeal checks if an existing deal can be updated
func TestUpdateDeal(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	updateJSON := `{"title": "Updated Deal"}`
	req, _ := http.NewRequest("PUT", "/v1/deals/1", bytes.NewBufferString(updateJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, true, response["success"])
	assert.Equal(t, "Updated Deal", response["data"].(map[string]interface{})["title"])
}

// TestGetDeal checks if a deal can be retrieved
func TestGetDeal(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/v1/deals/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, true, response["success"])
	assert.NotNil(t, response["data"])
}




// =======================
// Integration Tests
// =======================


func TestGetAllDealsIntegration(t *testing.T) {
	os.Setenv("PIPEDRIVE_API_KEY", "your_api_key")
	r := gin.Default()
	r.GET("/v1/deals", GetAllDeals)

	req, _ := http.NewRequest("GET", "/v1/deals", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var response api.GetAllDealsResponse
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.True(t, response.Success)
	fmt.Println("GetAllDeals response:", response)
}

func TestCreateDealIntegration(t *testing.T) {
	os.Setenv("PIPEDRIVE_API_KEY", "your_api_key")
	r := gin.Default()
	r.POST("/v1/deals", CreateDeal)

	dealRequest := api.CreateDealRequest{
		Title: "Test Deal",
	}
	body, _ := json.Marshal(dealRequest)
	req, _ := http.NewRequest("POST", "/v1/deals", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var response api.CreateDealResponse
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.True(t, response.Success)
	fmt.Println("CreateDeal response:", response)
}

func TestUpdateDealIntegration(t *testing.T) {
	os.Setenv("PIPEDRIVE_API_KEY", "your_api_key")
	r := gin.Default()
	r.PUT("/v1/deals/:id", UpdateDeal)

	dealRequest := api.CreateDealRequest{
		Title: "Test Deal Upd",
	}

	body, _ := json.Marshal(dealRequest)
	req, _ := http.NewRequest("PUT", "/v1/deals/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var response api.UpdateDealResponse
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.True(t, response.Success)
	fmt.Println("UpdateDeal response:", response)
}
