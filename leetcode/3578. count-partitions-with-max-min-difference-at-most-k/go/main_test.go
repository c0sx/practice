package leetcode3578

import "testing"

func TestCountPartitions(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{9, 4, 1, 3, 7}, 4, 6},
		{[]int{3, 3, 4}, 0, 2},
	}

	for _, tt := range tests {
		if got := countPartitions(tt.nums, tt.k); got != tt.want {
			t.Errorf("countPartitions(%v, %d) = %d; want %d", tt.nums, tt.k, got, tt.want)
		}
	}
}
