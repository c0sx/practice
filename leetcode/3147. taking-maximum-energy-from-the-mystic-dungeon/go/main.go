package leetcode3147

func maximumEnergy(energy []int, k int) int {
	n := len(energy)
	dp := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		if i+k >= n {
			dp[i] = energy[i]
		} else {
			dp[i] = energy[i] + dp[i+k]
		}
	}

	maxEnergy := dp[0]
	for i := 1; i < n; i++ {
		if dp[i] > maxEnergy {
			maxEnergy = dp[i]
		}
	}

	return maxEnergy
}
