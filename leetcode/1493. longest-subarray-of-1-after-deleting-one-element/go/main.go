package leetcode1493

func longestSubarray(nums []int) int {
	result := 0

	left := make([]int, len(nums))
	right := make([]int, len(nums))

	for i := 1; i < len(nums); i++ {
		prev := nums[i-1]

		if prev == 1 {
			left[i] = left[i-1] + 1
		}
	}

	for i := len(nums) - 2; i >= 0; i-- {
		next := nums[i+1]

		if next == 1 {
			right[i] = right[i+1] + 1
		}
	}

	for i := 0; i < len(nums); i++ {
		l := left[i]
		r := right[i]

		result = max(result, l+r)
	}

	return result
}
