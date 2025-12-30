package leetcode840

import "testing"

func TestNumMagicSquaresInside(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}, 1},
		{[][]int{{8}}, 0},
	}

	for _, tt := range tests {
		got := numMagicSquaresInside(tt.grid)
		if got != tt.want {
			t.Errorf("numMagicSquaresInside(%v) = %d; want %d", tt.grid, got, tt.want)
		}
	}
}
