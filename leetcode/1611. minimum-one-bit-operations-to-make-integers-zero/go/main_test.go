package leetcode1611

import "testing"

func TestMinimumOneBitOperations(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{3, 2},
		{6, 4},
	}

	for _, tt := range tests {
		if got := minimumOneBitOperations(tt.n); got != tt.want {
			t.Errorf("minimumOneBitOperations(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
