package leetcode1716

func totalMoney(n int) int {
	result := 0
	week := n / 7
	day := n % 7

	for i := 0; i < week; i++ {
		weekStart := 1 + i
		for j := 0; j < 7; j++ {
			result += weekStart + j
		}
	}

	weekStart := 1 + week
	for j := 0; j < day; j++ {
		result += weekStart + j
	}

	return result
}
