package leetcodeq2

import "testing"

func TestShuffle(t *testing.T) {
	tests := []struct {
		nums     []int
		n        int
		expected []int
	}{
		{[]int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}},
		{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{[]int{1, 1, 2, 2}, 2, []int{1, 2, 1, 2}},
	}

	for _, test := range tests {
		result := shuffle(test.nums, test.n)
		if !equal(result, test.expected) {
			t.Errorf("shuffle(%v, %d) = %v; expected %v", test.nums, test.n, result, test.expected)
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
