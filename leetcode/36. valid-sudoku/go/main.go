package leetcode36

func isValidSudoku(board [][]byte) bool {
	rows := len(board)
	cols := len(board[0])

	// check row
	for i := 0; i < rows; i++ {
		values := make(map[int]bool, 10)

		for j := 0; j < cols; j++ {
			if board[i][j] == '.' {
				continue
			}

			v := int(board[i][j] - '1')
			if values[v] {
				return false
			}

			values[v] = true
		}
	}

	// check col
	for j := 0; j < cols; j++ {
		values := make(map[int]bool, 10)

		for i := 0; i < rows; i++ {
			if board[i][j] == '.' {
				continue
			}

			v := int(board[i][j] - '1')
			if values[v] {
				return false
			}

			values[v] = true
		}
	}

	// check box
	for boxRow := 0; boxRow < 3; boxRow++ {
		for boxCol := 0; boxCol < 3; boxCol++ {
			values := make(map[int]bool, 10)

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					rowIdx := boxRow*3 + i
					colIdx := boxCol*3 + j

					if board[rowIdx][colIdx] == '.' {
						continue
					}

					v := int(board[rowIdx][colIdx] - '1')
					if values[v] {
						return false
					}

					values[v] = true
				}
			}
		}
	}

	return true
}
