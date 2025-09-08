package leetcode1317

func getNoZeroIntegers(n int) []int {
	for i := range n {
		b := n - i

		if hasZero(i) || hasZero(b) {
			continue
		}

		return []int{i, b}
	}

	return []int{}
}

func hasZero(n int) bool {
	if n == 0 {
		return true
	}

	for n > 0 {
		if n%10 == 0 {
			return true
		}

		n /= 10
	}

	return false
}
