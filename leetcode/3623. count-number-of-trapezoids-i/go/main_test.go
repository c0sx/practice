package leetcode3623

import "testing"

func TestCountTrapezoids(t *testing.T) {
	tests := []struct {
		points   [][]int
		expected int
	}{
		{[][]int{{1, 0}, {2, 0}, {3, 0}, {2, 2}, {3, 2}}, 3},
		{[][]int{{0, 0}, {1, 0}, {0, 1}, {2, 1}}, 1},
	}

	for _, test := range tests {
		result := countTrapezoids(test.points)
		if result != test.expected {
			t.Errorf("countTrapezoids(%v) = %d; want %d", test.points, result, test.expected)
		}
	}
}
