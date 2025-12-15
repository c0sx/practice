package leetcode2147

const MOD = 1e9 + 7

func numberOfWays(corridor string) int {
	ways := 1

	for i, n := 0, len(corridor); i < n; {
		seats := 0
		for ; i < n && seats != 2; i++ {
			if corridor[i] == 'S' {
				seats++
			}
		}

		if seats != 2 {
			return 0
		}

		plants := 0
		for ; i < n && corridor[i] == 'P'; i++ {
			plants++
		}

		if i < n {
			ways = (ways * (plants + 1)) % MOD
		}
	}

	return ways
}
