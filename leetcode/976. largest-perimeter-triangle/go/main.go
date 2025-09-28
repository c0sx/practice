package leetcode976

import (
	"sort"
)

func largestPerimeter(nums []int) int {
	sort.Ints(nums)

	n := len(nums)
	for i := n - 1; i >= 2; i-- {
		if nums[i-2]+nums[i-1] > nums[i] {
			p := nums[i] + nums[i-1] + nums[i-2]

			return p
		}
	}

	return 0
}
