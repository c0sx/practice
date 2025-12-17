package leetcode3573

import "testing"

/**
Example 1:

Input: prices = [1,7,9,8,2], k = 2

Output: 14

Example 2:

Input: prices = [12,16,19,19,8,1,19,13,9], k = 3

Output: 36

*/

func TestMaximumProfit(t *testing.T) {
	tests := []struct {
		prices []int
		k      int
		want   int64
	}{
		{[]int{1, 7, 9, 8, 2}, 2, 14},
		{[]int{12, 16, 19, 19, 8, 1, 19, 13, 9}, 3, 36},
	}

	for _, tt := range tests {
		if got := maximumProfit(tt.prices, tt.k); got != tt.want {
			t.Errorf("maximumProfit(%v, %d) = %d; want %d", tt.prices, tt.k, got, tt.want)
		}
	}
}
