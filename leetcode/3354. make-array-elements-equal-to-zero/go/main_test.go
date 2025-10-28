package leetcode3354

import "testing"

func TestCountValidSelections(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 0, 2, 0, 3}, 2},
		{[]int{2, 3, 4, 0, 4, 1, 0}, 0},
	}

	for _, tt := range tests {
		if got := countValidSelections(tt.nums); got != tt.want {
			t.Errorf("countValidSelections(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
