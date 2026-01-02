package leetcode961

import "testing"

func TestRepeatedNTimes(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 3}, 3},
		{[]int{2, 1, 2, 5, 3, 2}, 2},
		{[]int{5, 1, 5, 2, 5, 3, 5, 4}, 5},
	}

	for _, tt := range tests {
		if got := repeatedNTimes(tt.nums); got != tt.want {
			t.Errorf("repeatedNTimes(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
