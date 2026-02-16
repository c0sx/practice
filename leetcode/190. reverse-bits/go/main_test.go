package lc

import "testing"

func TestReverseBits(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{
			n:    43261596,
			want: 964176192,
		},
		{
			n:    2147483644,
			want: 1073741822,
		},
	}

	for _, tt := range tests {
		got := reverseBits(tt.n)
		if got != tt.want {
			t.Errorf("reverseBits(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
