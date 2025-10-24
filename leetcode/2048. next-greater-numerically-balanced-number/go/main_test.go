package leetcode2048

import "testing"

func TestNextBeautifulNumber(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 22},
		{1000, 1333},
		{3000, 3133},
	}

	for _, tt := range tests {
		if got := nextBeautifulNumber(tt.n); got != tt.want {
			t.Errorf("nextBeautifulNumber(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
