package leetcode3651

import "testing"

func TestMinCost(t *testing.T) {
	tests := []struct {
		grid [][]int
		k    int
		want int
	}{
		{
			grid: [][]int{{1, 3, 3}, {2, 5, 4}, {4, 3, 5}},
			k:    2,
			want: 7,
		},
		{
			grid: [][]int{{1, 2}, {2, 3}, {3, 4}},
			k:    1,
			want: 9,
		},
	}

	for _, tt := range tests {
		if got := minCost(tt.grid, tt.k); got != tt.want {
			t.Errorf("minCost(%v, %d) = %d; want %d", tt.grid, tt.k, got, tt.want)
		}
	}
}
