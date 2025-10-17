package leetcode3003

import "testing"

func TestMaxPartitionsAfterOperations(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want int
	}{
		{"accca", 2, 3},
		{"aabaab", 3, 1},
		{"xxyz", 1, 4},
		{"aabcacc", 2, 3},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maxPartitionsAfterOperations(tt.s, tt.k); got != tt.want {
				t.Errorf("maxPartitionsAfterOperations(%q, %d) = %d; want %d", tt.s, tt.k, got, tt.want)
			}
		})
	}
}
