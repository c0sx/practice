package leetcode1018

import "testing"

func TestPrefixesDivBy5(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []bool
	}{
		{
			nums:     []int{0, 1, 1},
			expected: []bool{true, false, false},
		},
		{
			nums:     []int{1, 1, 1},
			expected: []bool{false, false, false},
		},
		{
			nums:     []int{0, 1, 1, 1, 1, 1},
			expected: []bool{true, false, false, false, true, false},
		},
		{
			nums:     []int{1, 1, 1, 0, 1},
			expected: []bool{false, false, false, false, false},
		},
	}

	for _, test := range tests {
		result := prefixesDivBy5(test.nums)

		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("prefixesDivBy5(%v) = %v; expected %v", test.nums, result, test.expected)
				break
			}
		}
	}
}
