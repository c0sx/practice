package leetcode3074

import "testing"

func TestMinimumBoxes(t *testing.T) {
	tests := []struct {
		apple    []int
		capacity []int
		expected int
	}{
		{[]int{1, 3, 2}, []int{4, 3, 1, 5, 2}, 2},
		{[]int{5, 5, 5}, []int{2, 4, 2, 7}, 4},
		{[]int{1, 8, 3, 3, 5}, []int{3, 9, 5, 1, 9}, 3},
	}

	for _, test := range tests {
		if result := minimumBoxes(test.apple, test.capacity); result != test.expected {
			t.Errorf("For apple %v and capacity %v, expected %d but got %d", test.apple, test.capacity, test.expected, result)
		}
	}
}
