package leetcode3289

import "testing"

func TestGetSneakyNumbers(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{0, 1, 1, 0},
			expected: []int{0, 1},
		},
		{
			nums:     []int{0, 3, 2, 1, 3, 2},
			expected: []int{2, 3},
		},
		{
			nums:     []int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2},
			expected: []int{4, 5},
		},
	}

	for _, test := range tests {
		result := getSneakyNumbers(test.nums)
		if len(result) != len(test.expected) {
			t.Errorf("For input %v, expected length %d but got %d", test.nums, len(test.expected), len(result))
			continue
		}

		for i := range result {
			found := false
			for j := range test.expected {
				if result[i] == test.expected[j] {
					found = true
					break
				}
			}

			if !found {
				t.Errorf("For input %v, expected %v but got %v", test.nums, test.expected, result)
				break
			}
		}
	}
}
