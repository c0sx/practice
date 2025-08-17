package leetcode837

func new21Game(n int, k int, maxPts int) float64 {
	if k == 0 || n >= k+maxPts-1 {
		return 1.0
	}

	pp := make([]float64, n+1)
	pp[0] = 1.0
	sum := 1.0
	result := 0.0

	for i := 1; i <= n; i++ {
		pp[i] = sum / float64(maxPts)

		if i < k {
			sum += pp[i]
		} else {
			result += pp[i]
		}

		if i >= maxPts {
			sum -= pp[i-maxPts]
		}
	}

	return result
}
