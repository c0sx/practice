package leetcode3538

import "testing"

func TestSpecialTriplets(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{6, 3, 6}, 1},
		{[]int{0, 1, 0, 0}, 1},
		{[]int{8, 4, 2, 8, 4}, 2},
	}

	for _, tt := range tests {
		if got := specialTriplets(tt.nums); got != tt.want {
			t.Errorf("specialTriplets(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
