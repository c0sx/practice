package leetcode3577

const MOD = 1e9 + 7

func countPermutations(complexity []int) int {
	n := len(complexity)

	for i := 1; i < n; i++ {
		if complexity[i] <= complexity[0] {
			return 0
		}
	}

	result := 1
	for i := 2; i < n; i++ {
		result = (result * i) % MOD
	}

	return result
}
