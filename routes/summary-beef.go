package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/newbieDev-22/backend-challenge/controller"
)

// SummaryBeefRoutes sets up the routes for beef summary operations
func SummaryBeefRoutes(router *gin.Engine, summaryBeefController controller.SummaryBeefController) {
	// @Summary Get beef text summary
	// @Description Retrieves a summary of beef-related text including word count and statistics
	// @Tags beef
	// @Accept json
	// @Produce json
	// @Success 200 {object} model.SummaryBeef "Successfully retrieved beef summary"
	// @Failure 500 {object} gin.H "Internal Server Error"
	// @Router /beef/summary [get]
	router.GET("/beef/summary", func(c *gin.Context) {
		summaryBeef, err := summaryBeefController.GetSummaryBeef()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, summaryBeef)
	})
}
