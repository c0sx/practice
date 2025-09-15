package leetcode1935

func canBeTypedWords(text string, brokenLetters string) int {
	broken := make(map[rune]bool)
	for _, c := range brokenLetters {
		broken[c] = true
	}

	count := 0
	word := true
	for _, c := range text {
		if c == ' ' {
			if word {
				count++
			}

			word = true
			continue
		}

		if broken[c] {
			word = false
		}
	}

	if word {
		count++
	}

	return count
}
