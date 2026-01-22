package leetcode3507

func minimumPairRemoval(nums []int) int {
	result := 0

	for len(nums) > 1 {
		isAscending := true
		index := -1
		minSum := 1 << 30

		for i := 0; i < len(nums)-1; i++ {
			prev := nums[i]
			curr := nums[i+1]
			sum := prev + curr
			if prev > curr {
				isAscending = false
			}

			if sum < minSum {
				minSum = sum
				index = i
			}
		}

		if isAscending {
			break
		}

		result++
		nums[index] = minSum
		nums = append(nums[:index+1], nums[index+2:]...)
	}

	return result
}
