package leetcode3453

import "testing"

func TestSeparateSquares(t *testing.T) {
	tests := []struct {
		squares [][]int
		want    float64
	}{
		{
			squares: [][]int{{0, 0, 1}, {2, 2, 1}},
			want:    1.0,
		},
		{
			squares: [][]int{{0, 0, 2}, {1, 1, 1}},
			want:    1.16667,
		},
	}

	for _, tt := range tests {
		if got := separateSquares(tt.squares); got != tt.want {
			t.Errorf("separateSquares(%v) = %v, want %v", tt.squares, got, tt.want)
		}
	}
}
