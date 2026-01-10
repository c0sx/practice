package leetcode712

import "testing"

func TestMinimumDeleteSum(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected int
	}{
		{"sea", "eat", 231},
		{"delete", "leet", 403},
	}

	for _, test := range tests {
		if result := minimumDeleteSum(test.s1, test.s2); result != test.expected {
			t.Errorf("For s1: %s and s2: %s, expected %d but got %d", test.s1, test.s2, test.expected, result)
		}
	}
}
