package lc

func longestBalanced(s string) int {
	results := []int{
		solve1(s, 'a'),
		solve1(s, 'b'),
		solve1(s, 'c'),
		solve2(s, 'a', 'b'),
		solve2(s, 'a', 'c'),
		solve2(s, 'b', 'c'),
		solve3(s),
	}

	maxLen := 0
	for _, v := range results {
		if v > maxLen {
			maxLen = v
		}
	}

	return maxLen
}

func solve1(s string, a byte) int {
	result := 0
	count := 0

	for i := 0; i < len(s); i++ {
		c := s[i]

		if c == a {
			count++
			if count > result {
				result = count
			}
		} else {
			count = 0
		}
	}

	return result
}

func solve2(s string, a, b byte) int {
	result := 0
	aCounter := 0
	bCounter := 0

	prev := map[int]int{}
	for i := 0; i < len(s); i++ {
		c := s[i]

		if c != a && c != b {
			prev = map[int]int{}
			aCounter = 0
			bCounter = 0
			continue
		}

		if c == a {
			aCounter++
		} else {
			bCounter++
		}

		if aCounter == bCounter {
			v := aCounter * 2
			if v > result {
				result = v
			}
		} else {
			diff := aCounter - bCounter
			if pv, ok := prev[diff]; ok {
				v := i - pv
				if v > result {
					result = v
				}
			} else {
				prev[diff] = i
			}
		}
	}

	return result
}

func solve3(s string) int {
	result := 0
	aCounter, bCounter, cCounter := 0, 0, 0
	prev := map[int64]int{}

	for i := 0; i < len(s); i++ {
		c := s[i]

		switch c {
		case 'a':
			aCounter++
		case 'b':
			bCounter++
		case 'c':
			cCounter++
		}

		if aCounter == bCounter && bCounter == cCounter {
			result = i + 1
		} else {
			diff := int64(aCounter-bCounter)*100001 + int64(bCounter-cCounter)
			if pv, ok := prev[diff]; ok {
				v := i - pv
				if v > result {
					result = v
				}
			} else {
				prev[diff] = i
			}
		}
	}

	return result
}
