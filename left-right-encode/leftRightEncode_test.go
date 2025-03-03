package leftrightencode

import "testing"

func TestLeftRightDecode(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "LLRR=",
			input:    "LLRR=",
			expected: "210122",
		},
		{
			name:     "==RLL",
			input:    "==RLL",
			expected: "000210",
		},
		{
			name:     "=LLRR",
			input:    "=LLRR",
			expected: "221012",
		},
		{
			name:     "RRL=R",
			input:    "RRL=R",
			expected: "012001",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := LeftRightDecode(tc.input)
			if result != tc.expected {
				t.Errorf("LeftRightDecode(%q) = %q; want %q", tc.input, result, tc.expected)
			}
		})
	}
}
