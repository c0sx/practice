package leetcode3195

func minimumArea(grid [][]int) int {
	minRow := len(grid)
	maxRow := 0

	minCol := len(grid[0])
	maxCol := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
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
