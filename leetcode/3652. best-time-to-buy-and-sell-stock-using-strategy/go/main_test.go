package leetcode3652

import "testing"

/**
Example 1:

Input: prices = [4,2,8], strategy = [-1,0,1], k = 2

Output: 10

Example 2:

Input: prices = [5,4,3], strategy = [1,1,0], k = 2

Output: 9

*/

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices   []int
		strategy []int
		k        int
		want     int64
	}{
		{
			prices:   []int{4, 2, 8},
			strategy: []int{-1, 0, 1},
			k:        2,
			want:     10,
		},
		{
			prices:   []int{5, 4, 3},
			strategy: []int{1, 1, 0},
			k:        2,
			want:     9,
		},
	}

	for _, tt := range tests {
		if got := maxProfit(tt.prices, tt.strategy, tt.k); got != tt.want {
			t.Errorf("maxProfit() = %v, want %v", got, tt.want)
		}
	}
}
