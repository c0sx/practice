package leetcode1518

import "testing"

func TestNumWaterBottles(t *testing.T) {
	tests := []struct {
		numBottles  int
		numExchange int
		expected    int
	}{
		{9, 3, 13},
		{15, 4, 19},
	}

	for _, test := range tests {
		result := numWaterBottles(test.numBottles, test.numExchange)
		if result != test.expected {
			t.Errorf("numWaterBottles(%d, %d) = %d; want %d", test.numBottles, test.numExchange, result, test.expected)
		}
	}
}
