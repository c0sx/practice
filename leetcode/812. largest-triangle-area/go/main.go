package leetcode812

import (
	"math"
)

func largestTriangleArea(points [][]int) float64 {
	maxArea := 0.0
	for i := 0; i < len(points)-2; i++ {
		p1 := points[i]

		for j := i + 1; j < len(points)-1; j++ {
			p2 := points[j]

			for k := i + 2; k < len(points); k++ {
				p3 := points[k]

				area := calcArea(p1, p2, p3)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea / 2.0
}

func calcArea(p1, p2, p3 []int) float64 {
	x1 := p1[0]
	y1 := p1[1]
	x2 := p2[0]
	y2 := p2[1]
	x3 := p3[0]
	y3 := p3[1]

	v1x := x2 - x1
	v1y := y2 - y1
	v2x := x3 - x1
	v2y := y3 - y1

	crossProduct := v1x*v2y - v1y*v2x
	area := math.Abs(float64(crossProduct))

	return area
}
