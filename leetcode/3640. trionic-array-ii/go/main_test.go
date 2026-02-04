package lc

import "testing"

func TestMaxSumTrionic(t *testing.T) {
	tests := []struct {
		nums []int
		want int64
	}{
		{
			nums: []int{0, -2, -1, -3, 0, 2, -1},
			want: -4,
		},
		{
			nums: []int{1, 4, 2, 7},
			want: 14,
		},
		{
			nums: []int{1, 4, 2, 2, 3, 1, 2},
			want: 8,
		},
	}

	for _, tt := range tests {
		got := maxSumTrionic(tt.nums)
		if got != tt.want {
			t.Errorf("maxSumTrionic(%v) = %d, want = %d", tt.nums, got, tt.want)
		}
	}
}
