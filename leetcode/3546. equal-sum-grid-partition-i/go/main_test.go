package lc

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		grid [][]int
		want bool
	}{
		{
			grid: [][]int{{1, 4}, {2, 3}},
			want: true,
		},
		{
			grid: [][]int{{1, 3}, {2, 4}},
			want: false,
		},
		{
			grid: [][]int{{100000, 100000, 92687}},
			want: false,
		},
	}

	for _, tt := range tests {
		got := canPartitionGrid(tt.grid)
		if got != tt.want {
			t.Errorf("func(%v) = %v, want %v", tt.grid, got, tt.want)
		}
	}
}
