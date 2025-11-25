package leetcode1015

func smallestRepunitDivByK(k int) int {
	remainder := 0

	for length := 1; length <= k; length++ {
		remainder = (remainder*10 + 1) % k

		if remainder == 0 {
			return length
		}
	}

	return -1
}
