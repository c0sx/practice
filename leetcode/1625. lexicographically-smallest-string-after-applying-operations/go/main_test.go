package leetcode1625

import "testing"

func TestFindLexSmallestString(t *testing.T) {
	tests := []struct {
		s        string
		a, b     int
		expected string
	}{
		{"5525", 9, 2, "2050"},
		{"74", 5, 1, "24"},
		{"0011", 4, 2, "0011"},
	}

	for _, test := range tests {
		result := findLexSmallestString(test.s, test.a, test.b)
		if result != test.expected {
			t.Errorf("findLexSmallestString(%q, %d, %d) = %q; want %q", test.s, test.a, test.b, result, test.expected)
		}
	}
}
