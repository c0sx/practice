package lc

import (
	"sort"
)

func largestSubmatrix(matrix [][]int) int {
	result := 0
	rows := len(matrix)
	cols := len(matrix[0])
	prev := make([]int, cols)

	for i := 0; i < rows; i++ {
		curr := matrix[i]
		for j := 0; j < cols; j++ {
			if curr[j] > 0 {
				curr[j] += prev[j]
			}
		}

		sorted := make([]int, len(curr))
		copy(sorted, curr)
		sort.Sort(sort.Reverse(sort.IntSlice(sorted)))

		for i := 0; i < len(sorted); i++ {
			result = max(result, sorted[i]*(i+1))
		}

		prev = curr
	}

	return result
}
