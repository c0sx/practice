package leetcode3234

import "testing"

func TestNumberOfSubstrings(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"00011", 5},
		{"101101", 16},
	}

	for _, test := range tests {
		result := numberOfSubstrings(test.s)
		if result != test.expected {
			t.Errorf("countSubstringsWithDominantOnes(%q) = %d; want %d", test.s, result, test.expected)
		}
	}
}
