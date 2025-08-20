package leetcode1277

func countSquares(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	result := 0

	for i := range m {
		for j := range n {
			if matrix[i][j] == 0 {
				continue
			}

			if i > 0 && j > 0 {
				matrix[i][j] = min(matrix[i-1][j], matrix[i][j-1], matrix[i-1][j-1]) + 1
			}

			result += matrix[i][j]
		}
	}

	return result
}
