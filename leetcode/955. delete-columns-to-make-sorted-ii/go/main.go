package leetcode955

func minDeletionSize(strs []string) int {
	n := len(strs)
	m := len(strs[0])

	resolved := make([]bool, n-1)
	unresolved := n - 1
	result := 0

	for col := 0; col < m && unresolved > 0; col++ {
		needDelete := false

		for row := 0; row < n-1; row++ {
			if !resolved[row] && strs[row][col] > strs[row+1][col] {
				needDelete = true
				break
			}
		}

		if needDelete {
			result++
			continue
		}

		for row := 0; row < n-1; row++ {
			if !resolved[row] && strs[row][col] < strs[row+1][col] {
				resolved[row] = true
				unresolved--
			}
		}
	}

	return result
}
