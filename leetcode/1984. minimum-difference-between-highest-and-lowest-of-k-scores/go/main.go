package leetcode1984

import (
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	n := len(nums)
	minDiff := math.MaxInt
	sort.Ints(nums)

	for i := 0; i <= n-k; i++ {
		diff := abs(nums[i+k-1] - nums[i])
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
