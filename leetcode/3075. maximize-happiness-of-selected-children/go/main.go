package leetcode3075

import (
	"sort"
)

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})

	turns := 0
	var total int64
	for i := 0; i < k; i++ {
		total += int64(max(happiness[i]-turns, 0))
		turns++
	}

	return total
}

func max(a, b int) int {
	if a > b {
		return a
	}
	
	return b
}
