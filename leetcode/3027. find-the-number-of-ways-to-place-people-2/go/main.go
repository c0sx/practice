package leetcode3027

import (
	"math"
	"sort"
)

func numberOfPairs(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}

		return points[i][0] < points[j][0]
	})

	result := 0

	for i := range points {
		p1y := points[i][1]
		maxY := math.MinInt

		for j := i + 1; j < len(points); j++ {
			p2y := points[j][1]

			if maxY < p2y && p2y <= p1y {
				result++
				maxY = p2y
			}
		}
	}

	return result
}
