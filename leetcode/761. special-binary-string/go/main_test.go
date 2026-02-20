package lc

import "testing"

func TestMakeLargestSpecial(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			s:    "11011000",
			want: "11100100",
		},
		{
			s:    "10",
			want: "10",
		},
	}

	for _, tt := range tests {
		got := makeLargestSpecial(tt.s)
		if got != tt.want {
			t.Errorf("makeLargestSpecial(%s) = %s, want %s", tt.s, got, tt.want)
		}
	}
}
