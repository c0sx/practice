package leetcode3190

import "testing"

func TestMinimumOperations(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 2, 3, 4},
			want: 3,
		},
		{
			nums: []int{3, 6, 9},
			want: 0,
		},
	}

	for _, tt := range tests {
		if got := minimumOperations(tt.nums); got != tt.want {
			t.Errorf("minimumOperations(%v) = %v; want %v", tt.nums, got, tt.want)
		}
	}
}
