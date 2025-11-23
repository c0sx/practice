package leetcode1262

import "testing"

func TestMaxSumDivThree(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 6, 5, 1, 8}, 18},
		{[]int{4}, 0},
		{[]int{1, 2, 3, 4, 4}, 12},
	}

	for _, tt := range tests {
		if got := maxSumDivThree(tt.nums); got != tt.want {
			t.Errorf("maxSumDivThree(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
