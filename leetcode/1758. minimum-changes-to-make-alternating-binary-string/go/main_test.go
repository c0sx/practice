package lc

import "testing"

func TestMinOperations(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			s:    "0100",
			want: 1,
		},
		{
			s:    "10",
			want: 0,
		},
		{
			s:    "1111",
			want: 2,
		},
	}

	for _, tt := range tests {
		got := minOperations(tt.s)
		if got != tt.want {
			t.Errorf("minOperations(%s) = %d, want %d", tt.s, got, tt.want)
		}
	}
}
