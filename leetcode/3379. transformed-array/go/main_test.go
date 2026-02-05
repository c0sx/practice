package lc

import "testing"

func TestConstructTransformedArray(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{3, -2, 1, 1},
			want: []int{1, 1, 1, 3},
		},
		{
			nums: []int{-1, 4, -1},
			want: []int{-1, -1, 4},
		},
	}

	for _, tt := range tests {
		got := constructTransformedArray(tt.nums)
		if !equal(got, tt.want) {
			t.Errorf("constructTransformedArray(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
