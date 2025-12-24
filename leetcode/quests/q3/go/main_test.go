package leetcodeq3

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 1, 0, 1, 1, 1}, 3},
		{[]int{1, 0, 1, 1, 0, 1}, 2},
	}

	for _, test := range tests {
		result := findMaxConsecutiveOnes(test.nums)
		if result != test.expected {
			t.Errorf("findMaxConsecutiveOnes(%v) = %d; expected %d", test.nums, result, test.expected)
		}
	}
}
