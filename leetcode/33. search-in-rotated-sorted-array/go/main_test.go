package leetcode33

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
	}

	for _, tt := range tests {
		if got := search(tt.nums, tt.target); got != tt.want {
			t.Errorf("search(%v, %v) = %v; want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}
