package leetcode3510

import "testing"

func TestMinimumPairRemoval(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{5, 2, 3, 1},
			want: 2,
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 2},
			want: 0,
		},
	}

	for _, tt := range tests {
		if got := minimumPairRemoval(tt.nums); got != tt.want {
			t.Errorf("minimumPairRemoval(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
