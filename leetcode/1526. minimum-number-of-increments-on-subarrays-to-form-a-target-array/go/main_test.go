package leetcode1526

import "testing"

func TestMinNumberOperations(t *testing.T) {
	tests := []struct {
		target []int
		want   int
	}{
		{[]int{1, 2, 3, 2, 1}, 3},
		{[]int{3, 1, 1, 2}, 4},
		{[]int{3, 1, 5, 4, 2}, 7},
	}

	for _, tt := range tests {
		if got := minNumberOperations(tt.target); got != tt.want {
			t.Errorf("minNumberOperations(%v) = %d; want %d", tt.target, got, tt.want)
		}
	}
}
