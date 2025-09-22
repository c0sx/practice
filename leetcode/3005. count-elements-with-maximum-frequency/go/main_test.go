package leetcode3005

import "testing"

func TestMaxFrequencyElements(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 2, 3, 1, 4}, 4},
		{[]int{1, 2, 3, 4, 5}, 5},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maxFrequencyElements(tt.nums); got != tt.want {
				t.Errorf("maxFrequencyElements(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
