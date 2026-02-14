package lc

func champagneTower(poured int, query_row int, query_glass int) float64 {
	var a [102][102]float64
	a[0][0] = float64(poured)

	for r := 0; r <= query_row; r++ {
		for c := 0; c <= r; c++ {
			q := (a[r][c] - 1.0) / 2.0
			if q > 0 {
				a[r+1][c] += q
				a[r+1][c+1] += q
			}
		}
	}

	return min(1, a[query_row][query_glass])
}
