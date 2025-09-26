package leetcode611

import (
	"sort"
)

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	counter := 0

	for i := len(nums) - 1; i >= 2; i-- {
		l := 0
		r := i - 1

		for l < r {
			if nums[l]+nums[r] > nums[i] {
				counter += r - l
				r--
			} else {
				l++
			}
		}
	}

	return counter
}
