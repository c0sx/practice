package leetcode3459

import "testing"

func TestLenOfVDiagonal(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{2, 2, 1, 2, 2}, {2, 0, 2, 2, 0}, {2, 0, 1, 1, 0}, {1, 0, 2, 2, 2}, {2, 0, 0, 2, 2}}, 5},
		{[][]int{{2, 2, 2, 2, 2}, {2, 0, 2, 2, 0}, {2, 0, 1, 1, 0}, {1, 0, 2, 2, 2}, {2, 0, 0, 2, 2}}, 4},
		{[][]int{{1, 2, 2, 2, 2}, {2, 2, 2, 2, 0}, {2, 0, 0, 0, 0}, {0, 0, 2, 2, 2}, {2, 0, 0, 2, 0}}, 5},
		{[][]int{{1}}, 1},
	}

	for _, tt := range tests {
		if got := lenOfVDiagonal(tt.grid); got != tt.want {
			t.Errorf("lenOfVDiagonal(%v) = %v; want %v", tt.grid, got, tt.want)
		}
	}
}
