package leetcode3446

import (
	"sort"
)

func sortMatrix(grid [][]int) [][]int {
	rows := len(grid)

	for r := range rows {
		sortDiagonal(grid, r, 0, false)
	}

	for c := 1; c < rows; c++ {
		sortDiagonal(grid, 0, c, true)
	}

	return grid
}

func sortDiagonal(grid [][]int, startX, startY int, asc bool) {
	x := startX
	y := startY

	size := len(grid)
	items := make([]int, 0, size-x)

	for x < size && y < size {
		items = append(items, grid[x][y])
		x++
		y++
	}

	sort.Slice(items, func(i, j int) bool {
		if asc {
			return items[i] < items[j]
		}

		return items[i] > items[j]
	})

	x = startX
	y = startY

	i := 0
	for x < size && y < size {
		grid[x][y] = items[i]
		x++
		y++
		i++
	}
}
