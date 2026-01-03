package leetcode1411

import "testing"

func TestNumOfWays(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 12},
		{5000, 30228214},
	}

	for _, tt := range tests {
		if got := numOfWays(tt.n); got != tt.want {
			t.Errorf("numOfWays(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}
