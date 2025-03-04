package controller

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/newbieDev-22/backend-challenge/service"
)

// @title Backend Challenge API
// @version 1.0
// @description This is a backend service that provides beef-related text analysis
// @host localhost:8080
// @BasePath /api/v1

const baconIpsumAPI = "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"

type (
	// SummaryBeefController defines the interface for beef summary operations
	SummaryBeefController interface {
		// GetSummaryBeef godoc
		// @Summary Get beef text summary
		// @Description Fetches text from BaconIpsum API and analyzes beef-related word frequencies
		// @Tags beef
		// @Accept json
		// @Produce json
		// @Success 200 {object} model.SummaryBeef
		// @Failure 500 {object} gin.H "Internal Server Error with error message"
		// @Router /beef/summary [get]
		GetSummaryBeef(c *gin.Context)
	}

	summaryBeefControllerImpl struct {
		summaryBeefService service.SummaryBeefService
	}
)

// NewSummaryBeefController creates a new instance of SummaryBeefController
func NewSummaryBeefController(summaryBeefService service.SummaryBeefService) SummaryBeefController {
	return &summaryBeefControllerImpl{
		summaryBeefService: summaryBeefService,
	}
}

// HandleGetSummaryBeef handles the HTTP request for getting beef summary
func (c *summaryBeefControllerImpl) GetSummaryBeef(ctx *gin.Context) {
	// Get data from baconipsum
	response, err := http.Get(baconIpsumAPI)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to fetch data from API: %v", err)})
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to read response body: %v", err)})
		return
	}

	// Process text and get summary
	summaryBeef, err := c.summaryBeefService.GetSummaryBeef(string(body))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to process text summary: %v", err)})
		return
	}

	ctx.JSON(http.StatusOK, summaryBeef)
}
