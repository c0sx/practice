package leetcode3350

import "testing"

func TestMaxIncreasingSubarrays(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}, 3},
		{[]int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}, 2},
		{[]int{-8, 7, -16, -7, 18}, 2},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maxIncreasingSubarrays(tt.nums); got != tt.want {
				t.Errorf("maxIncreasingSubarrays(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
