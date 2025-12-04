package leetcode2211

import "testing"

func TestCountCollisions(t *testing.T) {
	tests := []struct {
		directions string
		want       int
	}{
		{"RLRSLL", 5},
		{"LLRR", 0},
	}

	for _, tt := range tests {
		if got := countCollisions(tt.directions); got != tt.want {
			t.Errorf("countCollisions(%v) = %v, want %v", tt.directions, got, tt.want)
		}
	}
}
