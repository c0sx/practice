package leetcode1493

import "testing"

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 0, 1}, 3},
		{[]int{0, 1, 0}, 1},
		{[]int{0, 1, 1, 1, 0, 1, 1, 0, 1}, 5},
		{[]int{1, 1, 1}, 2},
		{[]int{0, 0, 0}, 0},
	}

	for _, tt := range tests {
		if got := longestSubarray(tt.nums); got != tt.want {
			t.Errorf("longestSubarray(%v) = %v; want %v", tt.nums, got, tt.want)
		}
	}
}
