package leftrightencode

// LeftRightDecode converts a string of directional relationships into a sequence of numbers.
// Each character in the input string represents a relationship between adjacent numbers in the output:
//   - 'L': left number is greater than the right number
//   - 'R': left number is less than the right number
//   - '=': left number equals the right number
//
// The function returns a string of digits representing the smallest possible sequence
// that satisfies all the relationships. If the input contains invalid characters,
// an empty string is returned.

func LeftRightDecode(encoded string) string {
	if !isValidInput(encoded) {
		return ""
	}

	n := len(encoded)
	sequence := make([]int, n+1)

	// Forward pass: process R and = relationships
	for i := range encoded {
		switch encoded[i] {
		case 'R':
			sequence[i+1] = sequence[i] + 1
		case '=':
			sequence[i+1] = sequence[i]
		}
	}

	// Backward pass: process L relationships and ensure = relationships
	for i := n - 1; i >= 0; i-- {
		switch encoded[i] {
		case 'L':
			sequence[i] = max(sequence[i], sequence[i+1]+1)
		case '=':
			sequence[i] = sequence[i+1]
		}
	}

	return convertToString(sequence)
}

// isValidInput checks if the input string contains only valid characters (L, R, =)
func isValidInput(s string) bool {
	for _, ch := range s {
		if ch != 'L' && ch != 'R' && ch != '=' {
			return false
		}
	}
	return true
}

// convertToString converts a slice of integers to a string of digits
func convertToString(nums []int) string {
	result := make([]byte, len(nums))
	for i, num := range nums {
		result[i] = byte('0' + num)
	}
	return string(result)
}
