package leetcode2169

import "testing"

func TestCountOperations(t *testing.T) {
	tests := []struct {
		num1     int
		num2     int
		expected int
	}{
		{2, 3, 3},
		{10, 10, 1},
	}

	for _, test := range tests {
		if result := countOperations(test.num1, test.num2); result != test.expected {
			t.Errorf("countOperations(%d, %d) = %d; expected %d", test.num1, test.num2, result, test.expected)
		}
	}
}
