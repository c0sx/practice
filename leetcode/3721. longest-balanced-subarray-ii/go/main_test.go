package lc

import "testing"

func TestLongestBalanced(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{2, 5, 4, 3},
			want: 4,
		},
		{
			nums: []int{3, 2, 2, 5, 4},
			want: 5,
		},
		{
			nums: []int{1, 2, 3, 2},
			want: 3,
		},
	}

	for _, tt := range tests {
		got := longestBalanced(tt.nums)
		if got != tt.want {
			t.Errorf("longestBalanced(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}
