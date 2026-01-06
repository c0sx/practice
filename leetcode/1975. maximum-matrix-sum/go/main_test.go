package leetcode1975

import "testing"

func TestMaxMatrixSum(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   int64
	}{
		{[][]int{{1, -1}, {-1, 1}}, 4},
		{[][]int{{1, 2, 3}, {-1, -2, -3}, {1, 2, 3}}, 16},
	}

	for _, tt := range tests {
		if got := maxMatrixSum(tt.matrix); got != tt.want {
			t.Errorf("maxMatrixSum(%v) = %v, want %v", tt.matrix, got, tt.want)
		}
	}
}
