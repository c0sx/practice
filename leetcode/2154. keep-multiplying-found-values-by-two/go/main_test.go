package leetcode2154

import "testing"

func TestFindFinalValue(t *testing.T) {
	tests := []struct {
		nums     []int
		original int
		expected int
	}{
		{[]int{5, 3, 6, 1, 12}, 3, 24},
		{[]int{2, 7, 9}, 4, 4},
	}

	for _, tt := range tests {
		result := findFinalValue(tt.nums, tt.original)
		if result != tt.expected {
			t.Errorf("findFinalValue(%v, %d) = %d; expected %d", tt.nums, tt.original, result, tt.expected)
		}
	}
}
