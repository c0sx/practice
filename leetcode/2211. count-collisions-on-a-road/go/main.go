package leetcode2211

func countCollisions(directions string) int {
	count := 0
	flag := -1

	for _, dir := range directions {
		if dir == 'L' {
			if flag >= 0 {
				count += flag + 1
				flag = 0
			}
		} else if dir == 'S' {
			if flag > 0 {
				count += flag
			}

			flag = 0
		} else {
			if flag >= 0 {
				flag++
			} else {
				flag = 1
			}
		}
	}

	return count
}
