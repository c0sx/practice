package leetcode1262

func maxSumDivThree(nums []int) int {
	n := len(nums)
	dp := make([][3]int, n+1)
	for i := 0; i <= n; i++ {
		for j := 0; j < 3; j++ {
			dp[i][j] = -(1 << 31)
		}
	}

	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		num := nums[i-1]

		for j := 0; j < 3; j++ {
			exclude := dp[i-1][j]
			prev := (j - num%3 + 3) % 3
			include := dp[i-1][prev] + num

			dp[i][j] = max(exclude, include)
		}
	}

	return dp[n][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
