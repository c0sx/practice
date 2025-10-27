package leetcode2125

import "testing"

func TestNumberOfBeams(t *testing.T) {
	tests := []struct {
		bank     []string
		expected int
	}{
		{
			bank:     []string{"011001", "000000", "010100", "001000"},
			expected: 8,
		},
		{
			bank:     []string{"000", "111", "000"},
			expected: 0,
		},
	}

	for _, test := range tests {
		result := numberOfBeams(test.bank)
		if result != test.expected {
			t.Errorf("For bank %v, expected %d but got %d", test.bank, test.expected, result)
		}
	}
}
