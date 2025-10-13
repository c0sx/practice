package leetcode2273

func removeAnagrams(words []string) []string {
	for i := 1; i < len(words); i++ {
		a := words[i-1]
		b := words[i]

		if isAnagram(a, b) {
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	return words
}

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	count := make(map[rune]int)
	for _, ch := range a {
		count[ch]++
	}

	for _, ch := range b {
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}

	return true
}
