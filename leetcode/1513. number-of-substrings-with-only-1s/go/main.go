package leetcode1513

const MOD = 1_000_000_007

func numSub(s string) int {
	n := len(s)
	counter := 0
	total := 0

	for i := 0; i < n; i++ {
		if s[i] == '1' {
			counter++
		} else {
			total += counter * (counter + 1) / 2
			total %= MOD
			counter = 0
		}
	}

	total += counter * (counter + 1) / 2
	total %= MOD

	return total
}
