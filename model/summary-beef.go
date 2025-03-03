package model

// SummaryBeef represents the response structure for beef text analysis
// @Description Response model containing beef word frequency counts
type SummaryBeef struct {
	// Beef contains a map of beef-related words and their frequencies
	// @Description Map of beef-related words as keys and their frequencies as values
	Beef map[string]int `json:"beef" example:"{\"t-bone\":4,\"fatback\":1}"`
}
