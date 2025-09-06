package leetcode3495

import "testing"

func TestMinOperations(t *testing.T) {
	tests := []struct {
		queries [][]int
		want    int64
	}{
		{[][]int{{1, 2}, {2, 4}}, 3},
		{[][]int{{2, 6}}, 4},
	}

	for _, tt := range tests {
		if got := minOperations(tt.queries); got != tt.want {
			t.Errorf("minOperations(%v) = %v, want %v", tt.queries, got, tt.want)
		}
	}
}
