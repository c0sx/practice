package leetcode1611

func minimumOneBitOperations(n int) int {
	if n == 0 {
		return 0
	}

	bit := 0
	curr := 1
	for curr*2 <= n {
		curr *= 2
		bit++
	}

	return (1 << (bit + 1)) - 1 - minimumOneBitOperations(n^curr)
}
