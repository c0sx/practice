package q1

import "testing"

func TestFindErrorNums(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 2, 4}, []int{2, 3}},
		{[]int{1, 1}, []int{1, 2}},
		{[]int{2, 2}, []int{2, 1}},
		{[]int{3, 2, 2}, []int{2, 1}},
	}

	for _, test := range tests {
		if result := findErrorNums(test.nums); !equal(result, test.expected) {
			t.Errorf("For nums: %v, expected %v but got %v", test.nums, test.expected, result)
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
