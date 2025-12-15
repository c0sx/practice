package leetcode2110

import "testing"

func TestGetDescentPeriods(t *testing.T) {
	tests := []struct {
		prices []int
		want   int64
	}{
		{[]int{3, 2, 1, 4}, 7},
		{[]int{8, 6, 7, 7}, 4},
		{[]int{1}, 1},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := getDescentPeriods(tt.prices); got != tt.want {
				t.Errorf("getDescentPeriods(%v) = %v, want %v", tt.prices, got, tt.want)
			}
		})
	}
}
