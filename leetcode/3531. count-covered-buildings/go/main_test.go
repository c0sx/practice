package leetcode3531

import "testing"

func TestCuntCoveredBuildings(t *testing.T) {
	tests := []struct {
		n         int
		buildings [][]int
		expected  int
	}{
		{
			n:         3,
			buildings: [][]int{{1, 2}, {2, 2}, {3, 2}, {2, 1}, {2, 3}},
			expected:  1,
		},
		{
			n:         3,
			buildings: [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}},
			expected:  0,
		},
		{
			n:         5,
			buildings: [][]int{{1, 3}, {3, 2}, {3, 3}, {3, 5}, {5, 3}},
			expected:  1,
		},
		{
			n: 4,
			// [[2,4],[1,2],[3,1],[1,4],[2,3],[3,3],[2,2],[1,3]]
			buildings: [][]int{{2, 4}, {1, 2}, {3, 1}, {1, 4}, {2, 3}, {3, 3}, {2, 2}, {1, 3}},
			expected:  1,
		},
	}

	for _, test := range tests {
		if result := countCoveredBuildings(test.n, test.buildings); result != test.expected {
			t.Errorf("For n=%d and buildings=%v, expected %d but got %d", test.n, test.buildings, test.expected, result)
		}
	}
}
