package leetcode2257

import "testing"

func TestCountUnguardedCells(t *testing.T) {
	tests := []struct {
		m, n     int
		guards   [][]int
		walls    [][]int
		expected int
	}{
		{
			m: 4, n: 6,
			guards:   [][]int{{0, 0}, {1, 1}, {2, 3}},
			walls:    [][]int{{0, 1}, {2, 2}, {1, 4}},
			expected: 7,
		},
		{
			m: 3, n: 3,
			guards:   [][]int{{1, 1}},
			walls:    [][]int{{0, 1}, {1, 0}, {2, 1}, {1, 2}},
			expected: 4,
		},
		{
			m: 1, n: 10,
			guards:   [][]int{{0, 0}},
			walls:    [][]int{{0, 1}},
			expected: 8,
		},
		{
			m: 1, n: 100000,
			guards:   [][]int{{0, 0}},
			walls:    [][]int{{0, 1}},
			expected: 99998,
		},
	}

	for _, test := range tests {
		result := countUnguarded(test.m, test.n, test.guards, test.walls)
		if result != test.expected {
			t.Errorf("countUnguarded(%d, %d, %v, %v) = %d; want %d", test.m, test.n, test.guards, test.walls, result, test.expected)
		}
	}
}
