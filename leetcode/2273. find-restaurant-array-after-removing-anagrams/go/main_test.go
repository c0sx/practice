package leetcode2273

import "testing"

func TestRemoveAnagrams(t *testing.T) {
	tests := []struct {
		words    []string
		expected []string
	}{
		{[]string{"abba", "baba", "bbaa", "cd", "cd"}, []string{"abba", "cd"}},
		{[]string{"a", "b", "c", "d", "e"}, []string{"a", "b", "c", "d", "e"}},
	}

	for _, test := range tests {
		result := removeAnagrams(test.words)
		if !equal(result, test.expected) {
			t.Errorf("removeAnagrams(%v) = %v; want %v", test.words, result, test.expected)
		}
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
