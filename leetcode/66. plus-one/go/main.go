package leetcode66

func plusOne(digits []int) []int {
	n := len(digits)
	i := n - 1

	for i >= 0 {
		num := digits[i] + 1
		if num < 10 {
			digits[i] = num
			return digits
		}

		digits[i] = 0
		i--
	}

	return append([]int{1}, digits...)
}
