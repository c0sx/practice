package leetcode778

import "testing"

func TestSwimInWater(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{{0, 2}, {1, 3}},
			want: 3,
		},
		{
			grid: [][]int{
				{0, 1, 2, 3, 4},
				{24, 23, 22, 21, 5},
				{12, 13, 14, 15, 16},
				{11, 17, 18, 19, 20},
				{10, 9, 8, 7, 6},
			},
			want: 16,
		},
	}

	for _, tt := range tests {
		result := swimInWater(tt.grid)

		if result != tt.want {
			t.Errorf("swimInWater(%v) = %d, want %d", tt.grid, result, tt.want)
		}
	}
}
