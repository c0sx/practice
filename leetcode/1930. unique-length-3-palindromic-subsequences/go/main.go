package leetcode1930

func countPalindromicSubsequence(s string) int {
	letters := make(map[rune]bool)
	for _, ch := range s {
		letters[ch] = true
	}

	result := 0
	for letter := range letters {
		first, last := -1, 0

		for i, ch := range s {
			if ch == letter {
				if first == -1 {
					first = i
				}

				last = i
			}
		}

		between := make(map[byte]bool)
		for i := first + 1; i < last; i++ {
			between[s[i]] = true
		}

		result += len(between)
	}

	return result
}
