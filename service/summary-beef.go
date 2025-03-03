package service

import (
	"errors"
	"strings"

	"github.com/newbieDev-22/backend-challenge/model"
)

type (
	SummaryBeefService interface {
		GetSummaryBeef(beefData string) (model.SummaryBeef, error)
	}

	summaryBeefServiceImpl struct {
	}
)

func NewSummaryBeefService() SummaryBeefService {
	return &summaryBeefServiceImpl{}
}

type wordCountResult map[string]int

func processLine(line string) wordCountResult {
	localCount := make(wordCountResult)
	words := strings.Fields(line)
	for _, word := range words {
		if word = strings.TrimSpace(word); word != "" {
			localCount[word]++
		}
	}
	return localCount
}

func (s *summaryBeefServiceImpl) GetSummaryBeef(beefData string) (model.SummaryBeef, error) {
	if beefData == "" {
		return model.SummaryBeef{}, errors.New("beef data is empty")
	}

	// clean . and , to empty string
	beefData = strings.ReplaceAll(beefData, ".", "")
	beefData = strings.ReplaceAll(beefData, ",", "")

	// split beefData by newlines
	lines := strings.Split(strings.TrimSpace(beefData), "\n")

	// Create channels for results
	resultsChan := make(chan wordCountResult, len(lines))

	// Process each line concurrently
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		go func(text string) {
			resultsChan <- processLine(text)
		}(line)
	}

	// Combine results
	finalCount := make(wordCountResult)
	processedLines := 0
	expectedLines := 0

	// Count non-empty lines
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			expectedLines++
		}
	}

	// Collect results from all goroutines
	for result := range resultsChan {
		for word, count := range result {
			finalCount[word] += count
		}
		processedLines++
		if processedLines == expectedLines {
			close(resultsChan)
			break
		}
	}

	return model.SummaryBeef{
		Beef: finalCount,
	}, nil
}
