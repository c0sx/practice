package leetcode3025

import "testing"

func TestNumberOfPairs(t *testing.T) {
	tests := []struct {
		pairs [][]int
		want  int
	}{
		{[][]int{{1, 1}, {2, 2}, {3, 3}}, 0},
		{[][]int{{6, 2}, {4, 4}, {2, 6}}, 2},
		{[][]int{{3, 1}, {1, 3}, {1, 1}}, 2},
		{[][]int{{0, 3}, {6, 1}}, 1},
		{[][]int{{0, 0}, {0, 3}}, 1},
		{[][]int{{0, 0}, {2, 5}}, 0},
	}

	for _, test := range tests {
		result := numberOfPairs(test.pairs)
		if result != test.want {
			t.Errorf("numberOfPairs(%v) = %d; want %d", test.pairs, result, test.want)
		}
	}
}
