package leetcode2536

func rangeAddQueries(n int, queries [][]int) [][]int {
	diff := make([][]int, n+1)
	for i := range diff {
		diff[i] = make([]int, n+1)
	}

	for _, query := range queries {
		x1, y1, x2, y2 := query[0], query[1], query[2], query[3]

		diff[x1][y1]++
		diff[x1][y2+1]--
		diff[x2+1][y1]--
		diff[x2+1][y2+1]++
	}

	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x1 := 0
			if i > 0 {
				x1 = m[i-1][j]
			}

			x2 := 0
			if j > 0 {
				x2 = m[i][j-1]
			}

			x3 := 0
			if i > 0 && j > 0 {
				x3 = m[i-1][j-1]
			}

			m[i][j] = diff[i][j] + x1 + x2 - x3
		}
	}

	return m
}
