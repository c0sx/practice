package leetcode498

func findDiagonalOrder(mat [][]int) []int {
	rows := len(mat)
	cols := len(mat[0])
	result := make([]int, 0, rows*cols)
	isUp := true
	i := 0
	j := 0

	for len(result) < rows*cols {
		if isUp {
			for i >= 0 && j < cols {
				result = append(result, mat[i][j])
				i--
				j++
			}

			if i < 0 && j <= cols-1 {
				i = 0
			}

			if j == cols {
				i += 2
				j--
			}
		} else {
			for i < rows && j >= 0 {
				result = append(result, mat[i][j])
				i++
				j--
			}

			if i == rows {
				i--
				j += 2
			}

			if j < 0 {
				j = 0
			}
		}

		isUp = !isUp
	}

	return result
}
