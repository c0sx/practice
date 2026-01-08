package leetcode1458

import "testing"

func TestMaxDotProduct(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{
			nums1: []int{2, 1, -2, 5},
			nums2: []int{3, 0, -6},
			want:  18,
		},
		{
			nums1: []int{3, -2},
			nums2: []int{2, -6, 7},
			want:  21,
		},
		{
			nums1: []int{-1, -1},
			nums2: []int{1, 1},
			want:  -1,
		},
	}

	for _, tt := range tests {
		if got := maxDotProduct(tt.nums1, tt.nums2); got != tt.want {
			t.Errorf("maxDotProduct(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
		}
	}
}
