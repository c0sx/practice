package leetcode1039

import "testing"

func TestMinScoreTriangulation(t *testing.T) {
	tests := []struct {
		values []int
		want   int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{3, 7, 4, 5}, 144},
		{[]int{1, 3, 1, 4, 1, 5}, 13},
	}

	for _, tt := range tests {
		if got := minScoreTriangulation(tt.values); got != tt.want {
			t.Errorf("minScoreTriangulation(%v) = %v, want %v", tt.values, got, tt.want)
		}
	}
}
