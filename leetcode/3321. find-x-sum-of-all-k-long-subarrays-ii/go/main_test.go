package leetcode3321

import "testing"

func TestFindXSum(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		x    int
		want []int64
	}{
		{
			nums: []int{1, 1, 2, 2, 3, 4, 2, 3},
			k:    6,
			x:    2,
			want: []int64{6, 10, 12},
		},
		{
			nums: []int{3, 8, 7, 8, 7, 5},
			k:    2,
			x:    2,
			want: []int64{11, 15, 15, 15, 12},
		},
	}

	for _, tt := range tests {
		if got := findXSum(tt.nums, tt.k, tt.x); !equal(got, tt.want) {
			t.Errorf("findXSum(%v, %d, %d) = %v; want %v", tt.nums, tt.k, tt.x, got, tt.want)
		}
	}
}

func equal(a, b []int64) bool {
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
