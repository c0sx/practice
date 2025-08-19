package leetcode2348

import "testing"

func TestZeroFilledSubarray(t *testing.T) {
	tests := []struct {
		nums []int
		want int64
	}{
		{[]int{1, 3, 0, 0, 2, 0, 0, 4}, 6},
		{[]int{0, 0, 0, 2, 0, 0}, 9},
		{[]int{2, 10, 2019}, 0},
	}

	for _, tt := range tests {
		if got := zeroFilledSubarray(tt.nums); got != tt.want {
			t.Errorf("zeroFilledSubarray(%v) = %d; want %d", tt.nums, got, tt.want)
		}
	}
}
