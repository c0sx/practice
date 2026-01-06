package leetcode1975

import "math"

func maxMatrixSum(matrix [][]int) int64 {
	var sum int64 = 0
	negativeCount := 0
	minAbs := math.MaxInt64

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			value := matrix[i][j]
			sum += int64(abs(value))

			if value < 0 {
				negativeCount++
			}

			minAbs = min(minAbs, abs(value))
		}
	}

	if negativeCount%2 == 0 {
		return sum
	}

	return sum - int64(2*minAbs)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
