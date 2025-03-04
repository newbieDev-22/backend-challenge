package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/newbieDev-22/backend-challenge/controller"
)

// SummaryBeefRoutes sets up the routes for beef summary operations
func SummaryBeefRoutes(router *gin.Engine, summaryBeefController controller.SummaryBeefController) {
	// @Summary Get beef text summary
	// @Description Retrieves a summary of beef-related text by fetching content from BaconIpsum API and analyzing word frequencies
	// @Tags beef
	// @Accept json
	// @Produce json
	// @Success 200 {object} model.SummaryBeef "Successfully retrieved beef word frequency analysis"
	// @Failure 500 {object} gin.H "Internal Server Error with detailed error message"
	// @Router /beef/summary [get]
	router.GET("/beef/summary", summaryBeefController.GetSummaryBeef)
}
