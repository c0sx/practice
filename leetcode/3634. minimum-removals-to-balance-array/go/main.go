package lc

import "sort"

func minRemoval(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)

	result := n
	r := 0

	for l := 0; l < n; l++ {
		for r < n && nums[r] <= nums[l]*k {
			r++
		}

		curr := n - (r - l)
		if curr < result {
			result = curr
		}
	}

	return result
}
