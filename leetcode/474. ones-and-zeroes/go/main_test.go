package leetcode474

import "testing"

func TestFindMaxForm(t *testing.T) {
	tests := []struct {
		strs     []string
		m        int
		n        int
		expected int
	}{
		{
			strs:     []string{"10", "0001", "111001", "1", "0"},
			m:        5,
			n:        3,
			expected: 4,
		},
		{
			strs:     []string{"10", "0", "1"},
			m:        1,
			n:        1,
			expected: 2,
		},
	}

	for _, tt := range tests {
		result := findMaxForm(tt.strs, tt.m, tt.n)
		if result != tt.expected {
			t.Errorf("findMaxForm(%v, %d, %d) = %d; expected %d", tt.strs, tt.m, tt.n, result, tt.expected)
		}
	}
}
