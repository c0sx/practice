package leetcode1590

import "testing"

func TestMinSubarray(t *testing.T) {
	tests := []struct {
		nums     []int
		p        int
		expected int
	}{
		{[]int{3, 1, 4, 2}, 6, 1},
		{[]int{6, 3, 5, 2}, 9, 2},
		{[]int{1, 2, 3}, 3, 0},
	}

	for _, test := range tests {
		result := minSubarray(test.nums, test.p)
		if result != test.expected {
			t.Errorf("For nums %v and p %d, expected %d but got %d", test.nums, test.p, test.expected, result)
		}
	}
}
