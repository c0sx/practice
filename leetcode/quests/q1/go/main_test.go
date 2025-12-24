package leetcodeq1

import "testing"

func TestGetConcatenation(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 1}, []int{1, 2, 1, 1, 2, 1}},
		{[]int{1, 3, 2, 1}, []int{1, 3, 2, 1, 1, 3, 2, 1}},
	}

	for _, test := range tests {
		result := getConcatenation(test.nums)
		if !equal(result, test.expected) {
			t.Errorf("getConcatenation(%v) = %v; expected %v", test.nums, result, test.expected)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
