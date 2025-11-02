package leetcode2257

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	for _, wall := range walls {
		grid[wall[0]][wall[1]] = 1
	}

	for _, guard := range guards {
		grid[guard[0]][guard[1]] = 1
	}

	for _, guard := range guards {
		markGuardedCells(grid, guard[0], guard[1])
	}

	result := 0

	for _, row := range grid {
		for _, cell := range row {
			if cell == 0 {
				result++
			}
		}
	}

	return result
}

func markGuardedCells(grid [][]int, row, col int) {
	// up
	for r := row - 1; r >= 0; r-- {
		if grid[r][col] == 1 {
			break
		}

		grid[r][col] = 2
	}

	// down
	for r := row + 1; r < len(grid); r++ {
		if grid[r][col] == 1 {
			break
		}

		grid[r][col] = 2
	}

	// left
	for c := col - 1; c >= 0; c-- {
		if grid[row][c] == 1 {
			break
		}

		grid[row][c] = 2
	}

	// right
	for c := col + 1; c < len(grid[0]); c++ {
		if grid[row][c] == 1 {
			break
		}

		grid[row][c] = 2
	}
}
