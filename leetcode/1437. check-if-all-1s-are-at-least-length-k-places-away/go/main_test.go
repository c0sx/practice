package leetcode1437

import "testing"

func TestKLengthApart(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want bool
	}{
		{[]int{1, 0, 0, 0, 1, 0, 0, 1}, 2, true},
		{[]int{1, 0, 0, 1, 0, 1}, 2, false},
		{[]int{0, 0, 0, 1, 0, 1}, 2, false},
	}

	for _, tt := range tests {
		if got := kLengthApart(tt.nums, tt.k); got != tt.want {
			t.Errorf("kLengthApart(%v, %d) = %v; want %v", tt.nums, tt.k, got, tt.want)
		}
	}
}
