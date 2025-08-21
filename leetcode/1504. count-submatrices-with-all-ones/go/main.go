package leetcode1504

func numSubmat(mat [][]int) int {
	r := len(mat)
	c := len(mat[0])
	result := 0

	temp := make([][]int, r)
	for i := range temp {
		temp[i] = make([]int, c)
	}

	for i := range r {
		for j := range c {
			if mat[i][j] == 1 {
				if j == 0 {
					temp[i][j] = 1
				} else {
					temp[i][j] = temp[i][j-1] + 1
				}
			}
		}
	}

	for i := range r {
		for j := range c {
			m := temp[i][j]
			for k := i; k >= 0 && m > 0; k-- {
				m = min(m, temp[k][j])
				result += m
			}
		}
	}

	return result
}
