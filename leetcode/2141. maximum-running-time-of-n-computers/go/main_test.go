package leetcode2141

import "testing"

func TestMaxRunTime(t *testing.T) {
	tests := []struct {
		n         int
		batteries []int
		expected  int64
	}{
		{2, []int{3, 3, 3}, 4},
		{2, []int{1, 1, 1, 1}, 2},
	}

	for _, tt := range tests {
		result := maxRunTime(tt.n, tt.batteries)
		if result != tt.expected {
			t.Errorf("got %d, want %d", result, tt.expected)
		}
	}
}
