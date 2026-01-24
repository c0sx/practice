package leetcode1877

import (
	"sort"
)

func minPairSum(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	l := 0
	r := n - 1
	maxSum := 0

	for l < r {
		pairSum := nums[l] + nums[r]
		if pairSum > maxSum {
			maxSum = pairSum
		}

		l++
		r--
	}

	return maxSum
}
