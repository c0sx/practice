package leetcode3459


func lenOfVDiagonal(grid [][]int) int {
	dirs := [4][2]int{{-1, 1}, {1, 1}, {1, -1}, {-1, -1}}
	rows := len(grid)
	cols := len(grid[0])
	cache := make([][][4][2]int, rows)

	for i := range cache {
		cache[i] = make([][4][2]int, cols)
		for j := range cache[i] {
			for k := range cache[i][j] {
				for l := range cache[i][j][k] {
					cache[i][j][k][l] = -1
				}
			}
		}
	}

	var f func(i, j, dir int, turned bool, target int) int
	f = func(i, j, dir int, turned bool, target int) int {
		nx, ny := i+dirs[dir][0], j+dirs[dir][1]
		if nx < 0 || ny < 0 || nx >= rows || ny >= cols || grid[nx][ny] != target {
			return 0
		}

		turnInt := 0
		if turned {
			turnInt = 1
		}

		if cache[nx][ny][dir][turnInt] != -1 {
			return cache[nx][ny][dir][turnInt]
		}

		current := f(nx, ny, dir, turned, 2-target)
		if turned {
			current = max(current, f(nx, ny, (dir+1)%4, false, 2-target))
		}

		cache[nx][ny][dir][turnInt] = current + 1

		return current + 1
	}

	result := 0
	for i := range rows {
		for j := range cols {
			if grid[i][j] == 1 {
				for dir := 0; dir < 4; dir++ {
					result = max(result, f(i, j, dir, true, 2)+1)
				}
			}
		}
	}

	return result
}
