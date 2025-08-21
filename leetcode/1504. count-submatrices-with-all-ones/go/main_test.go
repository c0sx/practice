package leetcode1504

import "testing"

func TestNumSubmat(t *testing.T) {
	tests := []struct {
		mat  [][]int
		want int
	}{
		{
			mat: [][]int{
				{1, 0, 1},
				{1, 1, 0},
				{1, 1, 0},
			},
			want: 13,
		},
		{
			mat: [][]int{
				{0, 1, 1, 0},
				{0, 1, 1, 1},
				{1, 1, 1, 0},
			},
			want: 24,
		},
	}

	for _, tt := range tests {
		if got := numSubmat(tt.mat); got != tt.want {
			t.Errorf("numSubmat(%v) = %d; want %d", tt.mat, got, tt.want)
		}
	}
}
