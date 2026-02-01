package lc

import "testing"

func TestMinimumCost(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 2, 3, 12}, want: 6},
		{nums: []int{5, 4, 3}, want: 12},
		{nums: []int{10, 3, 1, 1}, want: 12},
	}

	for _, tt := range tests {
		got := minimumCost(tt.nums)
		if got != tt.want {
			t.Errorf("minimumCost(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
