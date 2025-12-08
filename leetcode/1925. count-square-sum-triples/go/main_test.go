package leetcode1925

import "testing"

func TestCountTriples(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{
			n:    5,
			want: 2,
		},
		{
			n:    10,
			want: 4,
		},
	}

	for _, tt := range tests {
		if got := countTriples(tt.n); got != tt.want {
			t.Errorf("countTriples(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
