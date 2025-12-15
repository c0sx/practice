package leetcode2147

import "testing"

func TestNumberOfWays(t *testing.T) {
	tests := []struct {
		corridor string
		want     int
	}{
		{"SSPPSPS", 3},
		{"PPSPSP", 1},
		{"S", 0},
	}

	for _, tt := range tests {
		if got := numberOfWays(tt.corridor); got != tt.want {
			t.Errorf("numberOfWays(%v) = %v, want %v", tt.corridor, got, tt.want)
		}
	}
}
