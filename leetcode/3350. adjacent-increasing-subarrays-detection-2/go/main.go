package leetcode3350

func maxIncreasingSubarrays(nums []int) int {
	n := len(nums)
	counter, prev := 1, 0
	result := 0

	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			counter++
		} else {
			prev = counter
			counter = 1
		}

		result = max(result, min(prev, counter))
		result = max(result, counter/2)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
