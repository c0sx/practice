package leetcode1039

func minScoreTriangulation(values []int) int {
	n := len(values)
	dp := make([][]int, n)

	for i := range n {
		dp[i] = make([]int, n)
	}

	var calc func(i, j int) int
	calc = func(i, j int) int {
		if i+1 >= j {
			return 0
		}

		if dp[i][j] != 0 {
			return dp[i][j]
		}

		minScore := 1<<31 - 1
		for k := i + 1; k < j; k++ {
			score := calc(i, k) + calc(k, j) + values[i]*values[k]*values[j]
			minScore = min(minScore, score)
		}

		dp[i][j] = minScore
		return minScore
	}

	return calc(0, n-1)
}
