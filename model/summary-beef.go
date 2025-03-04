package model

// SummaryBeef represents the response structure for beef text analysis
// @Description Response model containing beef word frequency analysis
type SummaryBeef struct {
	// Beef contains a map of beef-related words and their frequencies
	// @Description Map of beef-related words and their occurrence count in the text
	// @Example {"ribeye": 5, "brisket": 3, "tenderloin": 2}
	Beef map[string]int `json:"beef" example:"{\"ribeye\":5,\"brisket\":3,\"tenderloin\":2}" swaggertype:"object"`
}
