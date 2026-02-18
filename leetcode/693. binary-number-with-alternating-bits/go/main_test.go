package lc

import "testing"

func TestHasAlternatingBits(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{
			n:    5,
			want: true,
		},
		{
			n:    7,
			want: false,
		},
		{
			n:    11,
			want: false,
		},
	}

	for _, tt := range tests {
		got := hasAlternatingBits(tt.n)
		if got != tt.want {
			t.Errorf("hasAlternatingBits(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
