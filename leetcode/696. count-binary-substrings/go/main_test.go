package lc

import "testing"

func TestCountBinarySubstrings(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			s:    "00110011",
			want: 6,
		},
		{
			s:    "10101",
			want: 4,
		},
	}

	for _, tt := range tests {
		got := countBinarySubstrings(tt.s)
		if got != tt.want {
			t.Errorf("countBinarySubstrings(%s) = %d, want %d", tt.s, got, tt.want)
		}
	}

}
