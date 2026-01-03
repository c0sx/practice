package leetcode1411

const MOD int64 = 1000000007

func numOfWays(n int) int {
	var a, b int64 = 6, 6

	for i := 2; i <= n; i++ {
		newA := (2*a + 2*b) % MOD
		newB := (2*a + 3*b) % MOD
		a, b = newA, newB
	}

	return int((a + b) % MOD)
}
