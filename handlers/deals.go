package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

var apiToken = os.Getenv("PIPEDRIVE_API_KEY")
var pipedriveAPI = "https://api.pipedrive.com/v1/deals"

// Fetch all deals
func GetDeals(c *gin.Context) {
	client := resty.New()
	resp, err := client.R().
		SetQueryParam("api_token", apiToken).
		Get(pipedriveAPI)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deals": resp.String()})
}

// Create a new deal
func CreateDeal(c *gin.Context) {
	var deal map[string]interface{}
	if err := c.ShouldBindJSON(&deal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	client := resty.New()
	resp, err := client.R().
		SetQueryParam("api_token", apiToken).
		SetBody(deal).
		Post(pipedriveAPI)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": resp.String()})
}

// Update an existing deal
func UpdateDeal(c *gin.Context) {
	dealID := c.Query("id")
	if dealID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Deal ID is required"})
		return
	}

	var deal map[string]interface{}
	if err := c.ShouldBindJSON(&deal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	client := resty.New()
	resp, err := client.R().
		SetQueryParam("api_token", apiToken).
		SetBody(deal).
		Put(pipedriveAPI + "/" + dealID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": resp.String()})
}
