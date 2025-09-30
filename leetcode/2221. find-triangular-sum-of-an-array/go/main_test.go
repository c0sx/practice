package leetcode2221

import "testing"

func TestTriangularSum(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 8},
		{[]int{5}, 5},
	}

	for _, test := range tests {
		result := triangularSum(test.nums)
		if result != test.expected {
			t.Errorf("triangularSum(%v) = %d; want %d", test.nums, result, test.expected)
		}
	}
}
