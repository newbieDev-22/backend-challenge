package controller

import (
	"fmt"
	"io"
	"net/http"

	"github.com/newbieDev-22/backend-challenge/model"
	"github.com/newbieDev-22/backend-challenge/service"
)

const baconIpsumAPI = "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"

type (
	// SummaryBeefController defines the interface for beef summary operations
	SummaryBeefController interface {
		GetSummaryBeef() (model.SummaryBeef, error)
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

// GetSummaryBeef retrieves and processes beef-related text to generate statistics
func (c *summaryBeefControllerImpl) GetSummaryBeef() (model.SummaryBeef, error) {
	// Get data from baconipsum
	response, err := http.Get(baconIpsumAPI)
	if err != nil {
		return model.SummaryBeef{}, fmt.Errorf("failed to fetch data from API: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return model.SummaryBeef{}, fmt.Errorf("failed to read response body: %w", err)
	}

	// Process text and get summary
	summaryBeef, err := c.summaryBeefService.GetSummaryBeef(string(body))
	if err != nil {
		return model.SummaryBeef{}, fmt.Errorf("failed to process text summary: %w", err)
	}

	return summaryBeef, nil
}
