package lc

func finalPrices(prices []int) []int {
	n := len(prices)
	result := make([]int, n)
	stack := make([]int, 0, n)

	for i := range n {
		for len(stack) > 0 && prices[stack[len(stack)-1]] >= prices[i] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[idx] = prices[idx] - prices[i]
		}

		stack = append(stack, i)
	}

	for _, idx := range stack {
		result[idx] = prices[idx]
	}

	return result
}
