package leetcode3531

func countCoveredBuildings(n int, buildings [][]int) int {
	maxRow := make([]int, n+1)
	maxCol := make([]int, n+1)
	minRow := make([]int, n+1)
	minCol := make([]int, n+1)

	for i := range minRow {
		minRow[i] = n + 1
		minCol[i] = n + 1
	}

	for _, b := range buildings {
		x, y := b[0], b[1]

		maxRow[y] = max(maxRow[y], x)
		minRow[y] = min(minRow[y], x)
		maxCol[x] = max(maxCol[x], y)
		minCol[x] = min(minCol[x], y)
	}

	result := 0

	for _, b := range buildings {
		x, y := b[0], b[1]
		if x > minRow[y] && x < maxRow[y] && y > minCol[x] && y < maxCol[x] {
			result++
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
