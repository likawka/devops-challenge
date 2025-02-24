package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"devops-challenge/api"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

// GetAllDeals godoc
// @Summary Get all deals
// @Description Fetches all deals from Pipedrive
// @Tags deals
// @Accept json
// @Produce json
// @Param params query api.GetAllDealsParams false "Query parameters for filtering deals"
// @Success 200 {object} api.GetAllDealsResponse
// @Router /v1/deals [get]
func GetAllDeals(c *gin.Context) {
	var queryParams api.GetAllDealsParams
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

	rawDeals, ok := pipedriveResponse["data"].([]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected response format"})
		return
	}

	var deals []api.Deal
	for _, rawDeal := range rawDeals {
		dealBytes, _ := json.Marshal(rawDeal)
		var deal api.Deal
		if err := json.Unmarshal(dealBytes, &deal); err != nil {
			continue
		}
		deals = append(deals, deal)
	}

	c.JSON(http.StatusOK, api.GetAllDealsResponse{
		Success: true,
		Data:    deals,
	})
}
