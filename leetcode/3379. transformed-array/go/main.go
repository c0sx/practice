package lc

func constructTransformedArray(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	for i, num := range nums {
		var next int
		if num > 0 {
			next = (i + num) % n
		} else if num < 0 {
			next = (i + num%n + n) % n
		} else {
			next = i
		}

		result[i] = nums[next]
	}

	return result
}
