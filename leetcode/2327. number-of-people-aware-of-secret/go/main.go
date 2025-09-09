package leetcode2327

const mod = 1_000_000_007

func peopleAwareOfSecret(n int, delay int, forget int) int {
	people := make([]int, n)
	people[0] = 1
	counter := 0

	for i := 1; i < n; i++ {
		if i-delay >= 0 {
			counter = (counter + people[i-delay]) % mod
		}

		if i-forget >= 0 {
			counter = (counter - people[i-forget] + mod) % mod
		}

		people[i] = counter
	}

	total := 0
	for i := n - forget; i < n; i++ {
		if i >= 0 {
			total = (total + people[i]) % mod
		}
	}

	return total
}
