package leetcode3542

func minOperations(nums []int) int {
	stack := []int{}
	result := 0

	for _, num := range nums {
		for len(stack) > 0 && stack[len(stack)-1] > num {
			stack = stack[:len(stack)-1]
		}

		if num == 0 {
			continue
		}

		if len(stack) == 0 || stack[len(stack)-1] < num {
			stack = append(stack, num)
			result++
		}
	}

	return result
}
