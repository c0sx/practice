package leetcode1716

import "testing"

func TestTotalMoney(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{4, 10},
		{10, 37},
		{20, 96},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := totalMoney(tt.n); got != tt.want {
				t.Errorf("totalMoney(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
