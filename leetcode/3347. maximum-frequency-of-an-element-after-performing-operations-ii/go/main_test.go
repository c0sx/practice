package leetcode3347

import "testing"

func TestMaxFrequency(t *testing.T) {
	tests := []struct {
		nums          []int
		k             int
		numOperations int
		expected      int
	}{
		{
			nums:          []int{1, 4, 5},
			k:             1,
			numOperations: 2,
			expected:      2,
		},
		{
			nums:          []int{5, 11, 20, 20},
			k:             5,
			numOperations: 1,
			expected:      2,
		},
	}

	for _, test := range tests {
		result := maxFrequency(test.nums, test.k, test.numOperations)
		if result != test.expected {
			t.Errorf("maxFrequency(%v, %d, %d) = %d; expected %d", test.nums, test.k, test.numOperations, result, test.expected)
		}
	}
}
