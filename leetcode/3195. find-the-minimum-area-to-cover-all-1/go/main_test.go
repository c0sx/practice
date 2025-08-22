package leetcode3195

import "testing"

func TestMinimumArea(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{
				{0, 1, 0},
				{1, 0, 1},
			},
			want: 6,
		},
		// {
		// 	grid: [][]int{
		// 		{1, 0},
		// 		{0, 0},
		// 	},
		// 	want: 1,
		// },
		// {
		// 	grid: [][]int{
		// 		{0}, {1},
		// 	},
		// 	want: 1,
		// },
	}

	for _, tt := range tests {
		if got := minimumArea(tt.grid); got != tt.want {
			t.Errorf("minimumArea(%v) = %d; want %d", tt.grid, got, tt.want)
		}
	}
}
