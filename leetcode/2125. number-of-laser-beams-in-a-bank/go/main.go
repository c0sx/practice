package leetcode2125

func numberOfBeams(bank []string) int {
	prev, result := 0, 0

	for _, row := range bank {
		count := 0

		for _, ch := range row {
			if ch == '1' {
				count++
			}
		}

		if count > 0 {
			result += prev * count
			prev = count
		}
	}

	return result
}
