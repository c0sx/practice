package leetcode3542

import "testing"

func TestMinOperations(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{0, 2}, 1},
		{[]int{3, 1, 2, 1}, 3},
		{[]int{1, 2, 1, 2, 1, 2}, 4},
	}

	for _, tt := range tests {
		if got := minOperations(tt.nums); got != tt.want {
			t.Errorf("minOperations(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
