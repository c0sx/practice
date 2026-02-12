package lc

import "testing"

func TestLongestBalanced(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			s:    "abbac",
			want: 4,
		},
		{
			s:    "zzabccy",
			want: 4,
		},
		{
			s:    "aba",
			want: 2,
		},
	}

	for _, tt := range tests {
		got := longestBalanced(tt.s)
		if got != tt.want {
			t.Errorf("longestBalanced(%s) = %d, want %d", tt.s, got, tt.want)
		}
	}
}
