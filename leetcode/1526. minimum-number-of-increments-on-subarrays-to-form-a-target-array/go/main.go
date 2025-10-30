package leetcode1526

func minNumberOperations(target []int) int {
	n := len(target)
	counter := target[0]

	for i := 1; i < n; i++ {
		diff := target[i] - target[i-1]
		if diff > 0 {
			counter += diff
		}
	}

	return counter
}
