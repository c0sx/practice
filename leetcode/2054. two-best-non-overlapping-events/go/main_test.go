package leetcode2054

import "testing"

func TestMaxTwoEvents(t *testing.T) {
	tests := []struct {
		events [][]int
		want   int
	}{
		{[][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}}, 4},
		{[][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}}, 5},
		{[][]int{{1, 5, 3}, {1, 5, 1}, {6, 6, 5}}, 8},
	}

	for _, tt := range tests {
		if got := maxTwoEvents(tt.events); got != tt.want {
			t.Errorf("maxTwoEvents(%v) = %v; want %v", tt.events, got, tt.want)
		}
	}
}
