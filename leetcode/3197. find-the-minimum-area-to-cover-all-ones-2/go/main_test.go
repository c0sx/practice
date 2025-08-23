package leetcode3197

import "testing"

func TestMinimumSum(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{{1, 0, 1}, {1, 1, 1}},
			want: 5,
		},
		// {
		// 	grid: [][]int{{1, 0, 1, 0}, {0, 1, 0, 1}},
		// 	want: 5,
		// },
		// {
		// 	grid: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
		// 	want: 9,
		// },
	}
	for _, tt := range tests {
		if got := minimumSum(tt.grid); got != tt.want {
			t.Errorf("minimumSum(%v) = %v, want %v", tt.grid, got, tt.want)
		}
	}
}
