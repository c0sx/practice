package leetcode3021

import "math"

func flowerGame(n int, m int) int64 {
	nc := math.Ceil(float64(n) / 2)
	mc := math.Ceil(float64(m) / 2)

	nf := math.Floor(float64(n) / 2)
	mf := math.Floor(float64(m) / 2)

	return int64((nc * mf) + (nf * mc))
}
