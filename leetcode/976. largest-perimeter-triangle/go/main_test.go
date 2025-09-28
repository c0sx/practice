package leetcode976

import "testing"

func TestLargestPerimeter(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 1, 2}, 5},
		{[]int{1, 2, 1, 10}, 0},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := largestPerimeter(tt.nums); got != tt.want {
				t.Errorf("largestPerimeter(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
