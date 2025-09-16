package leetcode2197

import "testing"

func TestReplaceNonCoprimes(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{6, 4, 3, 2, 7, 6, 2}, []int{12, 7, 6}},
		{[]int{2, 2, 1, 1, 3, 3, 3}, []int{2, 1, 1, 3}},
	}

	for _, test := range tests {
		if got := replaceNonCoprimes(test.nums); !equal(got, test.want) {
			t.Errorf("replaceNonCoprimes(%v) = %v; want %v", test.nums, got, test.want)
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
