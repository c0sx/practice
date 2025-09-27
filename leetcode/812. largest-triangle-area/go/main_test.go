package leetcode812

import "testing"

func TestLargestTriangleArea(t *testing.T) {
	tests := []struct {
		points [][]int
		want   float64
	}{
		{[][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}, 2.0},
		{[][]int{{1, 0}, {0, 0}, {0, 1}}, 0.5},
		{[][]int{{9, 4}, {5, 8}, {6, 1}}, 12.0},
		{[][]int{{6, 3}, {5, 2}, {5, 8}, {0, 6}}, 15.0},
	}

	for _, tt := range tests {
		if got := largestTriangleArea(tt.points); got != tt.want {
			t.Errorf("largestTriangleArea(%v) = %v, want %v", tt.points, got, tt.want)
		}
	}
}
