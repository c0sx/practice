package leetcode3453

import "math"

func separateSquares(squares [][]int) float64 {
	maxY, totalArea := 0.0, 0.0
	for _, square := range squares {
		y, length := square[1], square[2]
		totalArea += float64(length * length)

		size := float64(y + length)
		if size > maxY {
			maxY = size
		}
	}

	var check func(limitY float64) bool
	check = func(limit_y float64) bool {
		area := 0.0

		for _, square := range squares {
			y, length := square[1], square[2]

			if float64(y) < limit_y {
				overlap := math.Min(limit_y-float64(y), float64(length))
				area += float64(length) * overlap
			}
		}

		return area >= totalArea/2.0
	}

	lo, hi := 0.0, maxY
	eps := 1e-5

	for math.Abs(hi-lo) > eps {
		mid := (hi + lo) / 2.0

		if check(mid) {
			hi = mid
		} else {
			lo = mid
		}
	}

	return math.Round(hi*1e5) / 1e5
}
