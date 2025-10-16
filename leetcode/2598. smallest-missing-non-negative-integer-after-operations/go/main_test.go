package leetcode2598

import (
	"testing"
)

func TestFindSmallestInteger(t *testing.T) {
	tests := []struct {
		nums  []int
		value int
		want  int
	}{
		{
			nums:  []int{1, -10, 7, 13, 6, 8},
			value: 5,
			want:  4,
		},
		{
			nums:  []int{1, -10, 7, 13, 6, 8},
			value: 7,
			want:  2,
		},
	}

	for _, tt := range tests {
		if got := findSmallestInteger(tt.nums, tt.value); got != tt.want {
			t.Errorf("findSmallestInteger(%v, %d) = %v, want %v", tt.nums, tt.value, got, tt.want)
		}
	}
}
