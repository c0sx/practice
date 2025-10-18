package leetcode3397

import "testing"

func TestMaxDistinctElements(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{1, 2, 2, 3, 3, 4}, 2, 6},
		{[]int{4, 4, 4, 4}, 1, 3},
	}

	for _, test := range tests {
		result := maxDistinctElements(test.nums, test.k)
		if result != test.expected {
			t.Errorf("maxDistinctElements(%v, %d) = %d; want %d", test.nums, test.k, result, test.expected)
		}
	}
}
