package maxpathlogic

import (
	"encoding/json"
	"os"
	"testing"
)

func TestFindMaxPath(t *testing.T) {

	testCases := []struct {
		name    string
		array2D [][]int
		want    int
	}{
		{
			name:    "empty triangle",
			array2D: [][]int{},
			want:    0,
		},
		{
			name: "small triangle",
			array2D: [][]int{
				{59},
				{73, 41},
				{52, 40, 53},
				{26, 53, 6, 34},
			},
			want: 237,
		},
		{
			name: "extreme case",
			array2D: [][]int{
				{1},
				{1, 100000},
				{1, -10000, 1},
			},
			want: 100002,
		},
	}

	// Test hard.json file
	hardData, err := os.ReadFile("../files/hard.json")
	if err != nil {
		t.Fatalf("Failed to read hard.json: %v", err)
	}

	var hardTriangle [][]int
	if err := json.Unmarshal(hardData, &hardTriangle); err != nil {
		t.Fatalf("Failed to unmarshal hard.json: %v", err)
	}

	testCases = append(testCases, struct {
		name    string
		array2D [][]int
		want    int
	}{
		name:    "hard triangle from file",
		array2D: hardTriangle,
		want:    7273,
	})

	// Run all test cases
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := FindMaxPath(testCase.array2D)
			if got != testCase.want {
				t.Errorf("FindMaxPath() = %v, want %v", got, testCase.want)
			}
		})
	}
}
