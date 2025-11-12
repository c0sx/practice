package leetcode2654

import "testing"

func TestMinOperations(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{2, 6, 3, 4},
			expected: 4,
		},
		{
			nums:     []int{2, 10, 6, 14},
			expected: -1,
		},
	}

	for _, test := range tests {
		if result := minOperations(test.nums); result != test.expected {
			t.Errorf("For nums=%v, expected %d but got %d", test.nums, test.expected, result)
		}
	}
}
