package leetcode1015

import "testing"

func TestSmallestRepunitDivByK(t *testing.T) {
	tests := []struct {
		k    int
		want int
	}{
		{k: 1, want: 1},
		{k: 2, want: -1},
		{k: 3, want: 3},
	}

	for _, tt := range tests {
		if got := smallestRepunitDivByK(tt.k); got != tt.want {
			t.Errorf("smallestRepunitDivByK(%d) = %d; want %d", tt.k, got, tt.want)
		}
	}
}
