package lc

import "testing"

func TestMinRemoval(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{
			nums: []int{2, 1, 5},
			k:    2,
			want: 1,
		},
		{
			nums: []int{1, 6, 2, 9},
			k:    3,
			want: 2,
		},
		{
			nums: []int{4, 6},
			k:    2,
			want: 0,
		},
	}

	for _, tt := range tests {
		got := minRemoval(tt.nums, tt.k)
		if got != tt.want {
			t.Errorf("minRemoval(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
		}
	}

}
