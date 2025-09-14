package leetcode966

import (
	"strings"
)

func spellchecker(wordlist []string, queries []string) []string {
	sensitive := make(map[string]string)
	insensitive := make(map[string]string)
	masked := make(map[string]string)

	for _, word := range wordlist {
		sensitive[word] = word

		lower := strings.ToLower(word)
		if _, ok := insensitive[lower]; !ok {
			insensitive[lower] = word
		}

		maskedWord := maskVowels(lower)
		if _, ok := masked[maskedWord]; !ok {
			masked[maskedWord] = word
		}
	}

	result := make([]string, len(queries))
	for i, query := range queries {
		if val, ok := sensitive[query]; ok {
			result[i] = val
			continue
		}

		lower := strings.ToLower(query)
		if val, ok := insensitive[lower]; ok {
			result[i] = val
			continue
		}

		maskedQuery := maskVowels(lower)
		if val, ok := masked[maskedQuery]; ok {
			result[i] = val
			continue
		}

		result[i] = ""
	}

	return result
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}

func maskVowels(word string) string {
	bs := []byte(word)
	for i := range bs {
		if isVowel(bs[i]) {
			bs[i] = '*'
		}
	}

	return string(bs)
}
