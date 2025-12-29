package leetcode756

import "testing"

func TestPyramidTransition(t *testing.T) {
	tests := []struct {
		bottom  string
		allowed []string
		want    bool
	}{
		{"BCD", []string{"BCC", "CDE", "CEA", "FFF"}, true},
		{"AAAA", []string{"AAB", "AAC", "BCD", "BBE", "DEF"}, false},
	}

	for _, tt := range tests {
		got := pyramidTransition(tt.bottom, tt.allowed)
		if got != tt.want {
			t.Errorf("pyramidTransition(%v, %v) = %v; want %v", tt.bottom, tt.allowed, got, tt.want)
		}
	}
}
