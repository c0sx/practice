package leetcode3047

import "testing"

func TestLargestSquareArea(t *testing.T) {
	tests := []struct {
		bottomLeft [][]int
		topRight   [][]int
		expected   int64
	}{
		{
			bottomLeft: [][]int{{1, 1}, {2, 2}, {3, 1}},
			topRight:   [][]int{{3, 3}, {4, 4}, {6, 6}},
			expected:   1,
		},
		{
			bottomLeft: [][]int{{0, 0}, {1, 3}, {1, 5}},
			topRight:   [][]int{{5, 5}, {5, 7}, {5, 9}},
			expected:   4,
		},
		{
			bottomLeft: [][]int{{1, 1}, {2, 2}, {1, 2}},
			topRight:   [][]int{{3, 3}, {4, 4}, {3, 4}},
			expected:   1,
		},
		{
			bottomLeft: [][]int{{1, 1}, {3, 3}, {3, 1}},
			topRight:   [][]int{{2, 2}, {4, 4}, {4, 2}},
			expected:   0,
		},
	}

	for _, test := range tests {
		result := largestSquareArea(test.bottomLeft, test.topRight)
		if result != test.expected {
			t.Errorf("For bottomLeft: %v and topRight: %v, expected %d but got %d",
				test.bottomLeft, test.topRight, test.expected, result)
		}
	}
}
