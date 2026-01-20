package leetcode3314

import "testing"

func TestMinBitwiseArray(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{2, 3, 5, 7},
			expected: []int{-1, 1, 4, 3},
		},
		{
			nums:     []int{11, 13, 31},
			expected: []int{9, 12, 15},
		},
	}

	for _, test := range tests {
		result := minBitwiseArray(test.nums)

		if !equals(result, test.expected) {
			t.Errorf("For nums %v, expected %v but got %v", test.nums, test.expected, result)
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
