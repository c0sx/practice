package leetcode3516

import "math"

func findClosest(x int, y int, z int) int {
	dx := math.Abs(float64(z - x))
	dy := math.Abs(float64(z - y))

	if dx < dy {
		return 1
	} else if dy < dx {
		return 2
	} else {
		return 0
	}
}
