package q2

func smallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if nums[j] < nums[i] {
				result[i]++
			}
		}
	}

	return result
}
