package leetcode3000

import "testing"

func TestAreaOfMaxDiagonal(t *testing.T) {
	tests := []struct {
		dimensions [][]int
		want       int
	}{
		{[][]int{{9, 3}, {8, 6}}, 48},
		{[][]int{{3, 4}, {4, 3}}, 12},
	}

	for _, tt := range tests {
		if got := areaOfMaxDiagonal(tt.dimensions); got != tt.want {
			t.Errorf("areaOfMaxDiagonal(%v) = %v, want %v", tt.dimensions, got, tt.want)
		}
	}
}
