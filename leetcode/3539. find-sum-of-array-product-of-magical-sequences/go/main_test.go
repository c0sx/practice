package leetcode3539

import "testing"

func TestMagicalSum(t *testing.T) {
	tests := []struct {
		m, k int
		nums []int
		want int
	}{
		{5, 5, []int{1, 10, 100, 10000, 1000000}, 991600007},
		{2, 2, []int{5, 4, 3, 2, 1}, 170},
		{1, 1, []int{28}, 28},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := magicalSum(tt.m, tt.k, tt.nums); got != tt.want {
				t.Errorf("magicalSum(%d, %d, %v) = %d; want %d", tt.m, tt.k, tt.nums, got, tt.want)
			}
		})
	}
}
