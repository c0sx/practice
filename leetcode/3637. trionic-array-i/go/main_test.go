package lc

import "testing"

func TestIsTrionic(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{nums: []int{1, 3, 5, 4, 2, 6}, want: true},
		{nums: []int{2, 1, 3}, want: false},
		{nums: []int{1, 2, 3, 2, 1, 2, 3}, want: true},
		{nums: []int{4, 1, 5, 2, 3}, want: false},
		{nums: []int{8, 9, 4, 6, 1}, want: false},
		{nums: []int{1, 2, 3, 2, 3, 4, 4}, want: false},
		{nums: []int{1, 1, 1, 9}, want: false},
		{nums: []int{9, 2, 9, 1}, want: false},
	}

	for _, tt := range tests {
		got := isTrionic(tt.nums)
		if got != tt.want {
			t.Errorf("isTrionic(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
