package leetcode3197

func minimumSum(grid [][]int) int {
	rotated := rotate(grid)

	return min(minimumArea(grid), minimumArea(rotated))
}

func rotate(grid [][]int) [][]int {
	n, m := len(grid), len(grid[0])
	ret := make([][]int, m)

	for i := range ret {
		ret[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ret[m-j-1][i] = grid[i][j]
		}
	}

	return ret
}

func minimumArea(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	result := n * m

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			s1 := minimumAreaInSubmatrix(grid, 0, i, 0, m-1)
			s2 := minimumAreaInSubmatrix(grid, i+1, n-1, 0, j)
			s3 := minimumAreaInSubmatrix(grid, i+1, n-1, j+1, m-1)

			result = min(result, s1+s2+s3)

			s4 := minimumAreaInSubmatrix(grid, 0, i, 0, j)
			s5 := minimumAreaInSubmatrix(grid, 0, i, j+1, m-1)
			s6 := minimumAreaInSubmatrix(grid, i+1, n-1, 0, m-1)

			result = min(result, s4+s5+s6)
		}
	}

	for i := 0; i+2 < n; i++ {
		for j := i + 1; j+1 < n; j++ {
			s1 := minimumAreaInSubmatrix(grid, 0, i, 0, m-1)
			s2 := minimumAreaInSubmatrix(grid, i+1, j, 0, m-1)
			s3 := minimumAreaInSubmatrix(grid, j+1, n-1, 0, m-1)

			result = min(result, s1+s2+s3)
		}
	}

	return result
}

func minimumAreaInSubmatrix(grid [][]int, u, d, l, r int) int {
	minRow := len(grid)
	maxRow := 0

	minCol := len(grid[0])
	maxCol := 0

	for i := u; i <= d; i++ {
		for j := l; j <= r; j++ {
			if grid[i][j] == 1 {
				if i < minRow {
					minRow = i
				}

				if i > maxRow {
					maxRow = i
				}

				if j < minCol {
					minCol = j
				}

				if j > maxCol {
					maxCol = j
				}
			}
		}
	}

	width := maxCol - minCol + 1
	height := maxRow - minRow + 1

	return width * height
}
