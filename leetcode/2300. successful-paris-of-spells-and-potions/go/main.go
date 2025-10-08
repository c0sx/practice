package leetcode2300

import (
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	result := make([]int, 0, len(spells))
	sort.Ints(potions)

	for _, spell := range spells {
		l := 0
		r := len(potions)

		for l < r {
			m := (l + r) / 2

			if int64(spell*potions[m]) >= success {
				r = m
			} else {
				l = m + 1
			}
		}

		result = append(result, len(potions)-l)
	}

	return result
}
