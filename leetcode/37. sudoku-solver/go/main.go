package leetcode37

const LEN = 9

func solveSudoku(board [][]byte) {
	rows := make([][]bool, LEN)
	cols := make([][]bool, LEN)
	boxes := make([][]bool, LEN)
	emptyCells := make([][2]int, 0)

	for i := range LEN {
		rows[i] = make([]bool, LEN)
		cols[i] = make([]bool, LEN)
		boxes[i] = make([]bool, LEN)
	}

	for i := range LEN {
		for j := range LEN {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				rows[i][num] = true
				cols[j][num] = true

				k := (i/3)*3 + (j / 3)
				boxes[k][num] = true
			} else {
				emptyCells = append(emptyCells, [2]int{i, j})
			}
		}
	}

	var backtrack func(i int) bool
	backtrack = func(i int) bool {
		if i == len(emptyCells) {
			return true
		}

		cell := emptyCells[i]
		row, col := cell[0], cell[1]
		box := (row/3)*3 + (col / 3)

		for c := byte('1'); c <= '9'; c++ {
			n := c - '1'
			if rows[row][n] || cols[col][n] || boxes[box][n] {
				continue
			}

			rows[row][n] = true
			cols[col][n] = true
			boxes[box][n] = true
			board[row][col] = c

			if backtrack(i + 1) {
				return true
			}

			rows[row][n] = false
			cols[col][n] = false
			boxes[box][n] = false
			board[row][col] = '.'
		}

		return false
	}

	backtrack(0)
}
