package leetcode2785

import (
	"sort"
)

func sortVowels(s string) string {
	bytes := []byte(s)
	vowels := make([]byte, 0)

	for _, b := range bytes {
		if isVowel(b) {
			vowels = append(vowels, b)
		}
	}

	sort.Slice(vowels, func(i, j int) bool {
		return vowels[i] < vowels[j]
	})

	j := 0
	for i := 0; i < len(s); i++ {
		if isVowel(bytes[i]) {
			bytes[i] = vowels[j]
			j++
		}
	}

	return string(bytes)
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' ||
		b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
}
