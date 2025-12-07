package leetcode1523

import "testing"

func TestCountOdds(t *testing.T) {
	tests := []struct {
		low      int
		high     int
		expected int
	}{
		{3, 7, 3},
		{8, 10, 1},
		{0, 10, 5},
	}

	for _, test := range tests {
		if output := countOdds(test.low, test.high); output != test.expected {
			t.Errorf("countOdds(%d, %d) = %d; expected %d", test.low, test.high, output, test.expected)
		}
	}
}
