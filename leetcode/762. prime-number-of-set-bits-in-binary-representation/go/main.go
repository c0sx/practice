package lc

import "math/bits"

func countPrimeSetBits(left int, right int) int {
	count := 0

	primes := map[int]bool{
		2:  true,
		3:  true,
		5:  true,
		7:  true,
		11: true,
		13: true,
		17: true,
		19: true,
	}

	for i := left; i <= right; i++ {
		b := bits.OnesCount(uint(i))
		if primes[b] {
			count++
		}
	}

	return count
}
