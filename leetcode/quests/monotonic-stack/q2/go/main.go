package lc

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	stack := make([]int, 0, n)

	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result[idx] = i - idx
		}

		stack = append(stack, i)
	}

	return result
}
