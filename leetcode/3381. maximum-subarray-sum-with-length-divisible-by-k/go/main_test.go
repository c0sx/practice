package leetcode3381

import "testing"

func TestMaxSubarraySum(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int64
	}{
		{[]int{1, 2}, 1, 3},
		{[]int{-1, -2, -3, -4, -5}, 4, -10},
		{[]int{-5, 1, 2, -3, 4}, 2, 4},
	}

	for _, tt := range tests {
		if got := maxSubarraySum(tt.nums, tt.k); got != tt.want {
			t.Errorf("maxSubarraySum(%v, %d) = %d; want %d", tt.nums, tt.k, got, tt.want)
		}
	}
}
