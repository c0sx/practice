package leetcode66

import "testing"

func TestPlusOne(t *testing.T) {
	tests := []struct {
		digits []int
		want   []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{9}, []int{1, 0}},
	}

	for _, tt := range tests {
		if got := plusOne(tt.digits); !equal(got, tt.want) {
			t.Errorf("plusOne(%v) = %v, want %v", tt.digits, got, tt.want)
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
