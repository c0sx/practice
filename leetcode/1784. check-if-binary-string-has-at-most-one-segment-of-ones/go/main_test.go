package lc

import "testing"

func TestCheckOneSegment(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{
			s:    "1001",
			want: false,
		},
		{
			s:    "110",
			want: true,
		},
		{
			s:    "1",
			want: true,
		},
	}

	for _, tt := range tests {
		got := checkOnesSegment(tt.s)
		if got != tt.want {
			t.Errorf("checkOnesSegment(%s) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
