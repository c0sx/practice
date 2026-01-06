package leetcode1390

import "testing"

func TestSumFourDivisors(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{21, 4, 7}, 32},
		{[]int{21, 21}, 64},
		{[]int{1, 2, 3, 4, 5}, 0},
	}

	for _, tt := range tests {
		if got := sumFourDivisors(tt.nums); got != tt.want {
			t.Errorf("sumFourDivisors(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
