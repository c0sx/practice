package leetcode1292

func maxSideLength(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])

	p := make([][]int, m+1)
	for i := range p {
		p[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			p[i][j] = p[i-1][j] + p[i][j-1] - p[i-1][j-1] + mat[i-1][j-1]
		}
	}

	l, r, result := 1, min(m, n), 0

	for l <= r {
		mid := (l + r) / 2
		find := false
		for i := 1; i <= m-mid+1; i++ {
			for j := 1; j <= n-mid+1; j++ {
				sum := p[i+mid-1][j+mid-1] - p[i-1][j+mid-1] - p[i+mid-1][j-1] + p[i-1][j-1]
				if sum <= threshold {
					find = true
					break
				}
			}

			if find {
				break
			}
		}

		if find {
			result = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
