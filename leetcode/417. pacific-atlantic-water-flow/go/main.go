package leetcode417

func pacificAtlantic(heights [][]int) [][]int {
	rows := len(heights)
	cols := len(heights[0])

	directions := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	counter := make([][]int, rows)
	visited := make([][]bool, rows)
	for i := range rows {
		counter[i] = make([]int, cols)
		visited[i] = make([]bool, cols)
	}

	var dfs func(int, int)
	dfs = func(r, c int) {
		if visited[r][c] {
			return
		}

		counter[r][c]++
		visited[r][c] = true

		height := heights[r][c]

		for _, dir := range directions {
			nextRow := r + dir[0]
			nextCol := c + dir[1]

			if nextRow >= 0 && nextRow < rows && nextCol >= 0 && nextCol < cols {
				if height <= heights[nextRow][nextCol] {
					dfs(nextRow, nextCol)
				}
			}
		}
	}

	for c := 0; c < cols; c++ {
		dfs(0, c)
	}

	for r := 0; r < rows; r++ {
		dfs(r, 0)
	}

	for i, row := range visited {
		for j := range row {
			visited[i][j] = false
		}
	}

	for c := 0; c < cols; c++ {
		dfs(rows-1, c)
	}

	for r := 0; r < rows; r++ {
		dfs(r, cols-1)
	}

	result := make([][]int, 0, rows)
	for i := range rows {
		for j := range cols {
			if counter[i][j] == 2 {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}
