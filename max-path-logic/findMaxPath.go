package maxpathlogic

// FindMaxPath finds the maximum path sum from top to bottom in a triangle of numbers.
// The path can only move to adjacent numbers in the row below.
//
// Parameters:
//   - array2D: A 2D slice representing a triangle of integers where each row i has i+1 elements.
//     The first row contains one element, the second row contains two elements, and so on.
//
// Returns:
//   - The maximum sum possible by following a path from the top to the bottom of the triangle.
//     Returns 0 if the input triangle is empty.
//
// Example:
//
//	triangle := [][]int{
//	    {3},
//	    {7, 4},
//	    {2, 4, 6},
//	}
//	maxSum := FindMaxPath(triangle)  // Returns 14 (path: 3 -> 7 -> 4)
//
// The function uses dynamic programming to build a table of maximum sums possible
// at each position, considering the maximum of the two possible paths from above.

func FindMaxPath(array2D [][]int) int {
	if len(array2D) == 0 {
		return 0
	}

	// Get the number of rows in the triangle
	rows := len(array2D)
	if rows == 0 {
		return 0
	}

	// Create DP table with the same size as the triangle
	dp := make([][]int, rows)
	for row := range rows {
		dp[row] = make([]int, row+1)
	}

	// Initialize the first row with the top value
	dp[0][0] = array2D[0][0]

	// Fill the DP table
	for row := 1; row < rows; row++ {
		for col := 0; col <= row; col++ {
			// Current value in the triangle
			current := array2D[row][col]
			lastRow := row - 1
			lastCol := col - 1

			if col == 0 {
				dp[row][col] = dp[lastRow][col] + current
				continue
			}

			// For rightmost element (j == i)
			if col == row {
				dp[row][col] = dp[lastRow][lastCol] + current
				continue
			}

			// For middle elements, take maximum of two possible paths
			dp[row][col] = Max(dp[lastRow][lastCol], dp[lastRow][col]) + current
		}
	}

	// Find maximum value in the last row
	maxSum := dp[rows-1][0]
	for col := 1; col < rows; col++ {
		maxSum = Max(maxSum, dp[rows-1][col])
	}

	return maxSum
}
