package leetcode1792

import "testing"

func TestMaxAverageRatio(t *testing.T) {
	tests := []struct {
		classes       [][]int
		extraStudents int
		expected      float64
	}{
		{
			[][]int{{1, 2}, {3, 5}, {2, 2}},
			2,
			0.78333,
		},
		{
			[][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}},
			4,
			0.53485,
		},
	}

	for _, test := range tests {
		if actual := maxAverageRatio(test.classes, test.extraStudents); actual != test.expected {
			t.Errorf("For %v and %v, expected %v but got %v", test.classes, test.extraStudents, test.expected, actual)
		}
	}
}
