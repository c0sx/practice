package leetcode3186

import "testing"

func TestMaximumTotalDamage(t *testing.T) {
	tests := []struct {
		power []int
		want  int64
	}{
		{[]int{1, 1, 3, 4}, 6},
		{[]int{7, 1, 6, 6}, 13},
		{[]int{7, 1, 6, 3}, 10},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maximumTotalDamage(tt.power); got != tt.want {
				t.Errorf("maximumTotalDamage(%v) = %d; want %d", tt.power, got, tt.want)
			}
		})
	}
}
