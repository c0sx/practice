package leetcode3349

func hasIncreasingSubarrays(nums []int, k int) bool {
	n := len(nums)

	prev := 0
	curr := 0
	s := 0

	for i := 0; i < n; i++ {
		curr++

		if i == n-1 || nums[i] >= nums[i+1] {
			s = max(s, curr/2, min(prev, curr))

			prev = curr
			curr = 0
		}
	}

	return s >= k
}

func max(a int, b ...int) int {
	m := a
	for _, v := range b {
		if v > m {
			m = v
		}
	}

	return m
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}
