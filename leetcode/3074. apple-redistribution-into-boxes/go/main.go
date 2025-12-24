package leetcode3074

import (
	"sort"
)

func minimumBoxes(apple []int, capacity []int) int {
	sum := 0
	for _, pack := range apple {
		sum += pack
	}

	sort.Slice(capacity, func(a, b int) bool {
		var capacityA = capacity[a]
		var capacityB = capacity[b]

		return capacityA > capacityB
	})

	n := len(capacity)
	count := 0

	for sum > 0 && count < n {
		num := capacity[count]
		count++
		sum -= num
	}

	return count
}
