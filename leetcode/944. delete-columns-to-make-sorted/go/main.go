package leetcode944

func minDeletionSize(strs []string) int {
	n := len(strs[0])
	s := len(strs)
	counter := 0

	for i := 0; i < n; i++ {
		for j := 1; j < s; j++ {
			prev := strs[j-1][i]
			curr := strs[j][i]

			if curr < prev {
				counter++
				break
			}
		}
	}

	return counter
}
