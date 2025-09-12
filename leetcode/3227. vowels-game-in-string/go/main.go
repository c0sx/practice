package leetcode3227

func doesAliceWin(s string) bool {
	for _, r := range s {
		if isVowel(r) {
			return true
		}
	}

	return false
}

func isVowel(b rune) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}
