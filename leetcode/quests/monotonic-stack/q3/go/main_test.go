package lc

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		heights []int
		want    int
	}{
		{
			heights: []int{2, 1, 5, 6, 2, 3},
			want:    10,
		},
		{
			heights: []int{2, 4},
			want:    4,
		},
		{
			heights: []int{1},
			want:    1,
		},
		{
			heights: []int{9, 0},
			want:    9,
		},
		{
			heights: []int{2, 0, 2},
			want:    2,
		},
		{
			heights: []int{1, 1},
			want:    2,
		},
	}

	for _, tt := range tests {
		got := largestRectangleArea(tt.heights)
		if got != tt.want {
			t.Errorf("func(%v) = %d, want %d", tt.heights, got, tt.want)
		}
	}
}
