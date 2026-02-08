package lc

import "testing"

func TestMinimumDeletions(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			s:    "aababbab",
			want: 2,
		},
		{
			s:    "bbaaaaabb",
			want: 2,
		},
	}

	for _, tt := range tests {
		got := minimumDeletions(tt.s)
		if got != tt.want {
			t.Errorf("minimumDeletions(%v) = %d, want %d", tt.s, got, tt.want)
		}
	}
}
