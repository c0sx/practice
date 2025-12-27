package leetcode2402

import "testing"

func TestMostBooked(t *testing.T) {
	tests := []struct {
		n        int
		meetings [][]int
		expected int
	}{
		{
			n:        2,
			meetings: [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}},
			expected: 0,
		},
		{
			n:        3,
			meetings: [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}},
			expected: 1,
		},
	}

	for _, test := range tests {
		result := mostBooked(test.n, test.meetings)
		if result != test.expected {
			t.Errorf("mostBooked(%d, %v) = %d; expected %d", test.n, test.meetings, result, test.expected)
		}
	}
}
