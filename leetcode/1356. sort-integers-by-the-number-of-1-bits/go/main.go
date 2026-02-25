package lc

import (
	"math/bits"
	"sort"
)

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		a := arr[i]
		b := arr[j]
		aOnes := bits.OnesCount(uint(a))
		bOnes := bits.OnesCount(uint(b))

		if aOnes == bOnes {
			return a < b
		}

		return aOnes < bOnes
	})

	return arr
}
