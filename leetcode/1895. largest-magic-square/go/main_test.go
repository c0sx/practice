package leetcode1895

import "testing"

func TestLargestMagicSquare(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{7, 1, 4, 5, 6}, {2, 5, 1, 6, 4}, {1, 5, 4, 3, 2}, {1, 2, 7, 3, 4}}, 3},
		{[][]int{{5, 1, 3, 1}, {9, 3, 3, 1}, {1, 3, 3, 8}}, 2},
	}

	for _, tt := range tests {
		if got := largestMagicSquare(tt.grid); got != tt.want {
			t.Errorf("largestMagicSquare(%v) = %v, want %v", tt.grid, got, tt.want)
		}
	}
}
