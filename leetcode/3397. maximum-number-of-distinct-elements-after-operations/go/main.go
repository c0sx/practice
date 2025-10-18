package leetcode3397

import (
	"math"
	"sort"
)

func maxDistinctElements(nums []int, k int) int {
	sort.Ints(nums)
	counter := 0
	prev := math.MinInt32

	for _, num := range nums {
		current := min(max(num-k, prev+1), num+k)
		if current > prev {
			counter++
			prev = current
		}
	}

	return counter
}
