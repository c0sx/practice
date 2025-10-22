package leetcode3347

import (
	"sort"
)

func maxFrequency(nums []int, k int, numOperations int) int {
	sort.Ints(nums)
	freq := calculateFreq(nums)
	maxFreq := 1
	n := len(nums)

	l, r := 0, 0
	for i := 0; i < n; i++ {
		for ; r < n && nums[r]-k <= nums[i]; r++ {
		}

		for ; l < i && nums[i]-k > nums[l]; l++ {
		}

		total := r - l
		maxFreq = max(maxFreq, freq[i]+min(numOperations, total-freq[i]))
	}

	l = 0
	for r = 0; r < n; r++ {
		for ; l < r && (r-l+1 > numOperations || nums[r]-nums[l] > 2*k); l++ {
		}

		if l <= r {
			maxFreq = max(maxFreq, r-l+1)
		}
	}

	return maxFreq
}

func calculateFreq(nums []int) []int {
	n := len(nums)
	freq := make([]int, n)

	l, r := 0, 0
	for {
		for ; r < n && nums[r] == nums[l]; r++ {
		}

		for i := l; i < r; i++ {
			freq[i] = r - l
		}

		if r >= n {
			break
		}

		l = r
	}

	return freq
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
