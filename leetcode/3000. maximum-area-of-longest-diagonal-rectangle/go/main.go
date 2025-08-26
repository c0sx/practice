package leetcode3000

import "math"

func areaOfMaxDiagonal(dimensions [][]int) int {
	max := 0.0
	maxArea := 0

	for _, d := range dimensions {
		a := float64(d[0])
		b := float64(d[1])
		c := math.Sqrt(a*a + b*b)
		area := d[0] * d[1]

		if c > max {
			max = c
			maxArea = area
		} else if c == max && area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
