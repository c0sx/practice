package lc

func minimumDeletions(s string) int {
	count := 0
	deletions := 0

	for _, c := range s {
		if c == 'b' {
			count++
		} else {
			deletions = min(deletions+1, count)
		}
	}

	return deletions
}
