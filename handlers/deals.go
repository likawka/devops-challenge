package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"devops-challenge/api"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

// GetDeals godoc
// @Summary Get all deals
// @Description Returns all deals from Pipedrive
// @Tags deals
// @Accept json
// @Produce json
// @Param params query api.GetDealsParams false "Query parameters for filtering deals"
// @Success 200 {object} api.GetDealsResponse
// @Router /deals [get]
func GetDeals(c *gin.Context) {
	var queryParams api.GetDealsParams
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	client := resty.New()
	resp, err := client.R().
		SetQueryParam("api_token", os.Getenv("PIPEDRIVE_API_KEY")).
		SetQueryParamsFromValues(c.Request.URL.Query()).
		Get("https://api.pipedrive.com/v1/deals")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var pipedriveResponse map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &pipedriveResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	formattedResponse := api.GetDealsResponse{
		Success:        true,
		Data:           pipedriveResponse["data"],
		AdditionalData: pipedriveResponse["additional_data"],
	}

	c.JSON(http.StatusOK, formattedResponse)
}

// CreateDeal godoc
// @Summary Add a new deal
// @Description Creates a new deal in Pipedrive using query parameters
// @Tags deals
// @Accept json
// @Produce json
// @Param params query api.CreateDealParams true "Deal data to be created"
// @Success 200 {object} api.CreateDealResponse
// @Router /deals [post]
func CreateDeal(c *gin.Context) {
	var dealParams api.CreateDealParams
	if err := c.ShouldBindQuery(&dealParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParam("api_token", os.Getenv("PIPEDRIVE_API_KEY")).
		SetQueryParamsFromValues(c.Request.URL.Query()). // Automatically pass all query params
		Post("https://api.pipedrive.com/v1/deals")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var pipedriveResponse map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &pipedriveResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	formattedResponse := api.CreateDealResponse{
		Success: true,
		Data:    pipedriveResponse["data"],
	}

	c.JSON(http.StatusOK, formattedResponse)
}

