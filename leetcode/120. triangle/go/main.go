package leetcode120

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}

	tmp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			tmp[j] = min(tmp[j], tmp[j+1]) + triangle[i][j]
		}
	}

	return tmp[0]
}
