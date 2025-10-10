package leetcode3147

import "testing"

func TestMaximumEnergy(t *testing.T) {
	tests := []struct {
		energy []int
		k      int
		want   int
	}{
		{[]int{5, 2, -10, -5, 1}, 3, 3},
		{[]int{-2, -3, -1}, 2, -1},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maximumEnergy(tt.energy, tt.k); got != tt.want {
				t.Errorf("maximumEnergy(%v, %d) = %d; want %d", tt.energy, tt.k, got, tt.want)
			}
		})
	}
}
