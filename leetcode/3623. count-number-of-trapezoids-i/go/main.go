package leetcode3623

const MOD = 1e9 + 7

func countTrapezoids(points [][]int) int {
	groups := make(map[int]int)

	for _, point := range points {
		y := point[1]

		groups[y]++
	}

	sum := 0
	result := 0

	for _, count := range groups {
		edge := (count * (count - 1)) / 2

		result = (result + sum*edge) % MOD
		sum = (sum + edge) % MOD
	}

	return result
}
