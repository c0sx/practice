package leetcode1277

import "testing"

func TestCountSquares(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   int
	}{
		{
			matrix: [][]int{
				{0, 1, 1, 1},
				{1, 1, 1, 1},
				{0, 1, 1, 1},
			},
			want: 15,
		},
		{
			matrix: [][]int{
				{1, 0, 1},
				{1, 1, 0},
				{1, 1, 0},
			},
			want: 7,
		},
	}

	for _, tt := range tests {
		if got := countSquares(tt.matrix); got != tt.want {
			t.Errorf("countSquares(%v) = %d; want %d", tt.matrix, got, tt.want)
		}
	}
}
