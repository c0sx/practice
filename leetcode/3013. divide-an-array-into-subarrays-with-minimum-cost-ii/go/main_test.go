package lc

import "testing"

func TestMinimumCost(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		dist int
		want int64
	}{
		{
			nums: []int{1, 3, 2, 6, 4, 2},
			k:    3,
			dist: 3,
			want: 5,
		},
		{
			nums: []int{10, 1, 2, 2, 2, 1},
			k:    4,
			dist: 3,
			want: 15,
		},
		{
			nums: []int{10, 8, 18, 9},
			k:    3,
			dist: 1,
			want: 36,
		},
	}

	for _, tt := range tests {
		got := minimumCost(tt.nums, tt.k, tt.dist)
		if got != tt.want {
			t.Errorf("minimumCost(%v, %d, %d) = %d, want %d", tt.nums, tt.k, tt.dist, got, tt.want)
		}
	}
}
