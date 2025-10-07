package leetcode1488

func avoidFlood(rains []int) []int {
	n := len(rains)
	lakes := []int{}
	rainDays := make(map[int]int)

	result := make([]int, n)
	for i := range n {
		result[i] = 1
	}

	li := 0
	for i, rain := range rains {
		result[i] = 1
		if rain == 0 {
			lakes = append(lakes, i)
		} else {
			result[i] = -1

			day, ok := rainDays[rain]
			if ok {
				di := -1
				for fi, fp := range lakes[li:] {
					if fp > day {
						di = li + fi
						break
					}
				}

				if di == -1 {
					return []int{}
				}

				result[lakes[di]] = rain
				lakes[di] = -1

				if di == li {
					li++
				}
			}

			rainDays[rain] = i
		}
	}

	return result
}
