package q2

import "testing"

func TestSmallerNumbersThanCurrent(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{8, 1, 2, 2, 3},
			expected: []int{4, 0, 1, 1, 3},
		},
		{
			nums:     []int{6, 5, 4, 8},
			expected: []int{2, 1, 0, 3},
		},
		{
			nums:     []int{7, 7, 7, 7},
			expected: []int{0, 0, 0, 0},
		},
	}

	for _, test := range tests {
		result := smallerNumbersThanCurrent(test.nums)

		if !equals(result, test.expected) {
			t.Errorf("For input %v, expected %v but got %v", test.nums, test.expected, result)
			continue
		}
	}
}

func equals(a, b []int) bool {
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
