package leetcode3446

import "testing"

func TestSortMatrix(t *testing.T) {
	tests := []struct {
		grid     [][]int
		expected [][]int
	}{
		{
			grid:     [][]int{{1, 7, 3}, {9, 8, 2}, {4, 5, 6}},
			expected: [][]int{{8, 2, 3}, {9, 6, 7}, {4, 5, 1}},
		},
		{
			grid:     [][]int{{0, 1}, {1, 2}},
			expected: [][]int{{2, 1}, {1, 0}},
		},
		{
			grid:     [][]int{{1}},
			expected: [][]int{{1}},
		},
	}

	for _, test := range tests {
		if output := sortMatrix(test.grid); !equal(output, test.expected) {
			t.Errorf("sortMatrix(%v) = %v; expected %v", test.grid, output, test.expected)
		}
	}
}

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
