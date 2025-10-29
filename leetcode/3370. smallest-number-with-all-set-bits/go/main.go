package leetcode3370

func smallestNumber(n int) int {
	x := 1
	for x <= n {
		x <<= 1
	}

	return x - 1
}
