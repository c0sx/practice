package leetcode3461

import "testing"

func TestHasSameDigits(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"3902", true},
		{"34789", false},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := hasSameDigits(tt.s); got != tt.want {
				t.Errorf("hasSameDigits(%q) = %v; want %v", tt.s, got, tt.want)
			}
		})
	}
}
