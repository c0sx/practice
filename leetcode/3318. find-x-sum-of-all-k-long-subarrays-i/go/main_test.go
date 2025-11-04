package leetcode3318

import "testing"

func TestFindXSum(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		x    int
		want []int
	}{
		{
			nums: []int{1, 1, 2, 2, 3, 4, 2, 3},
			k:    6,
			x:    2,
			want: []int{6, 10, 12},
		},
		{
			nums: []int{3, 8, 7, 8, 7, 5},
			k:    2,
			x:    2,
			want: []int{11, 15, 15, 15, 12},
		},
	}

	for _, tt := range tests {
		got := findXSum(tt.nums, tt.k, tt.x)
		if !equal(got, tt.want) {
			t.Errorf("findXSum(%v, %d, %d) = %v; want %v", tt.nums, tt.k, tt.x, got, tt.want)
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
