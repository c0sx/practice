package lc

func minimumCost(nums []int) int {
	first := nums[0]
	second := 100
	third := 100

	for i := 1; i < len(nums); i++ {
		if nums[i] < second {
			third = second
			second = nums[i]
		} else if nums[i] < third {
			third = nums[i]
		}
	}

	return first + second + third
}
