package leetcode2483

func bestClosingTime(customers string) int {
	n := len(customers)
	penalty := 0
	min := 0
	result := n

	for i := n - 1; i >= 0; i-- {
		if customers[i] == 'Y' {
			penalty++
		} else {
			penalty--
		}

		if penalty <= min {
			min = penalty
			result = i
		}
	}

	return result
}
