package leetcode3370

import "testing"

func TestSmallestNumber(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{5, 7},
		{10, 15},
		{3, 3},
	}

	for _, test := range tests {
		result := smallestNumber(test.n)
		if result != test.expected {
			t.Errorf("smallestNumber(%d) = %d; want %d", test.n, result, test.expected)
		}
	}
}
