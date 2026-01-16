package leetcode2975

import "testing"

func TestMaximizeSquareArea(t *testing.T) {
	tests := []struct {
		m        int
		n        int
		hFences  []int
		vFences  []int
		expected int
	}{
		{
			m:        4,
			n:        3,
			hFences:  []int{2, 3},
			vFences:  []int{2},
			expected: 4,
		},
		{
			m:        6,
			n:        7,
			hFences:  []int{2},
			vFences:  []int{4},
			expected: -1,
		},
	}

	for _, test := range tests {
		result := maximizeSquareArea(test.m, test.n, test.hFences, test.vFences)
		if result != test.expected {
			t.Errorf("maximizeSquareArea(%d, %d, %v, %v) = %d; expected %d",
				test.m, test.n, test.hFences, test.vFences, result, test.expected)
		}
	}
}
