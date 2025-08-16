package leetcode1323

func maximum69Number(num int) int {
	var digits []int

	for num > 0 {
		digits = append([]int{num % 10}, digits...)
		num /= 10
	}

	for i, digit := range digits {
		if digit == 6 {
			digits[i] = 9
			break
		}
	}

	result := 0

	for _, digit := range digits {
		result = result*10 + digit
	}

	return result
}
