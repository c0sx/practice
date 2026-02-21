package lc

import "testing"

func TestCountPrimeSetBits(t *testing.T) {
	tests := []struct {
		left  int
		right int
		want  int
	}{
		{
			left:  6,
			right: 10,
			want:  4,
		},
		{
			left:  10,
			right: 15,
			want:  5,
		},
	}

	for _, tt := range tests {
		got := countPrimeSetBits(tt.left, tt.right)
		if got != tt.want {
			t.Errorf("countPrimeSetBits(%d, %d) = %d, want %d", tt.left, tt.right, got, tt.want)
		}
	}
}
