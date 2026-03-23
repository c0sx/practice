package lc

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{{-1, -2, -3}, {-2, -3, -3}, {-3, -3, -2}},
			want: -1,
		},
		{
			grid: [][]int{{1, -2, 1}, {1, -2, 1}, {3, -4, 1}},
			want: 8,
		},
		{
			grid: [][]int{{1, 3}, {0, -4}},
			want: 0,
		},
	}

	for _, tt := range tests {
		got := maxProductPath(tt.grid)
		if got != tt.want {
			t.Errorf("fn(%v) = %d, want %d", tt.grid, got, tt.want)
		}
	}
}
