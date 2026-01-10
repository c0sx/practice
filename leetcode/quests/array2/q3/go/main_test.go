package q3

import "testing"

func TestFindDisappearedNumbers(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{4, 3, 2, 7, 8, 2, 3, 1},
			expected: []int{5, 6},
		},
		{
			nums:     []int{1, 1},
			expected: []int{2},
		},
	}

	for _, test := range tests {
		result := findDisappearedNumbers(test.nums)

		if !equal(result, test.expected) {
			t.Errorf("For input %v, expected %v but got %v", test.nums, test.expected, result)
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
