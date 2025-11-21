package leetcode1930

import "testing"

func TestCountPalindromicSubsequence(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{
			s:        "aabca",
			expected: 3,
		},
		{
			s:        "adc",
			expected: 0,
		},
		{
			s:        "bbcbaba",
			expected: 4,
		},
	}

	for _, test := range tests {
		if output := countPalindromicSubsequence(test.s); output != test.expected {
			t.Errorf("For input %v, expected %v but got %v", test.s, test.expected, output)
		}
	}
}
