package leetcode1304

func sumZero(n int) []int {
	result := make([]int, n)

	for i := 0; i < n/2; i++ {
		result[i] = -(i + 1)
		result[n-i-1] = i + 1
	}

	return result
}
