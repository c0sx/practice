package leetcode2749

import "math/bits"

func makeTheIntegerZero(num1 int, num2 int) int {
	counter := 0

	for {
		counter += 1

		target := num1 - counter*num2
		if target < 0 {
			break
		}

		b := bits.OnesCount(uint(target))
		if b <= counter && counter <= target {
			return counter
		}
	}

	return -1
}
