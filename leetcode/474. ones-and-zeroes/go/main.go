package leetcode474

func findMaxForm(strs []string, m int, n int) int {
	size := len(strs)

	dp := make([][][]int, size+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)

		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}

	countZerosOnes := func(s string) (int, int) {
		zeros := 0

		for _, c := range s {
			if c == '0' {
				zeros++
			}
		}

		ones := len(s) - zeros

		return zeros, ones
	}

	for i := 1; i <= size; i++ {
		zeros, ones := countZerosOnes(strs[i-1])

		for availableZeros := 0; availableZeros <= m; availableZeros++ {
			for availableOnes := 0; availableOnes <= n; availableOnes++ {
				dp[i][availableZeros][availableOnes] = dp[i-1][availableZeros][availableOnes]

				if availableZeros >= zeros && availableOnes >= ones {
					dp[i][availableZeros][availableOnes] = max(
						dp[i][availableZeros][availableOnes],
						dp[i-1][availableZeros-zeros][availableOnes-ones]+1,
					)
				}
			}
		}
	}

	return dp[size][m][n]
}
