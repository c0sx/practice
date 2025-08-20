package leetcode1277

func countSquares(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	result := 0

	tmp := make([][]int, m)
	for i := range tmp {
		tmp[i] = make([]int, n)
	}

	for i := range m {
		for j := range n {
			if i == 0 || j == 0 {
				tmp[i][j] = matrix[i][j]
			} else if matrix[i][j] == 1 {
				tmp[i][j] = min(tmp[i-1][j], tmp[i][j-1], tmp[i-1][j-1]) + 1
			}

			result += tmp[i][j]
		}
	}

	return result
}
