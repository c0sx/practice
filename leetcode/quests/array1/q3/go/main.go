package leetcodeq3

func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	n := len(nums)
	currentCount := 0

	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			currentCount++
		} else {
			if currentCount > maxCount {
				maxCount = currentCount
			}

			currentCount = 0
		}
	}

	if currentCount > maxCount {
		maxCount = currentCount
	}

	return maxCount
}
