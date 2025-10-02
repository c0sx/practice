package leetcode3100

import "testing"

func TestMaxBottlesDrunk(t *testing.T) {
	tests := []struct {
		numBottles  int
		numExchange int
		want        int
	}{
		{13, 6, 15},
		{10, 3, 13},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maxBottlesDrunk(tt.numBottles, tt.numExchange); got != tt.want {
				t.Errorf("maxBottlesDrunk(%d, %d) = %d, want %d", tt.numBottles, tt.numExchange, got, tt.want)
			}
		})
	}
}
