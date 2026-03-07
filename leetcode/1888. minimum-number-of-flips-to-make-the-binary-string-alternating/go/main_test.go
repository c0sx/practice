package lc

import "testing"

func TestMinFlips(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			s:    "111000",
			want: 2,
		},
		{
			s:    "010",
			want: 0,
		},
	}

	for _, tt := range tests {
		got := minFlips(tt.s)
		if got != tt.want {
			t.Errorf("minFlips(%s) = %d, want %d", tt.s, got, tt.want)
		}
	}
}
