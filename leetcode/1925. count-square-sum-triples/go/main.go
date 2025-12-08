package leetcode1925

import "math"

func countTriples(n int) int {
	count := 0

	for a := 1; a <= n; a++ {
		for b := a; b <= n; b++ {
			candidate := math.Sqrt(math.Pow(float64(a), 2) + math.Pow(float64(b), 2))
			if candidate > float64(n) {
				break
			}

			if candidate == math.Floor(candidate) {
				count += 2
			}
		}
	}

	return count
}
