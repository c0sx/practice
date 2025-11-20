package leetcode757

import "testing"

func TestIntersectionSizeTwo(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      int
	}{
		{
			[][]int{{1, 3}, {3, 7}, {8, 9}},
			5,
		},
		{
			[][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}},
			3,
		},
		{
			[][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}},
			5,
		},
	}

	for _, tt := range tests {
		if got := intersectionSizeTwo(tt.intervals); got != tt.want {
			t.Errorf("intersectionSizeTwo(%v) = %v, want %v", tt.intervals, got, tt.want)
		}
	}
}
