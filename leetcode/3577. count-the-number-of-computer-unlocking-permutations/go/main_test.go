package leetcode3577

import "testing"

func TestCountPermutations(t *testing.T) {
	tests := []struct {
		complexity []int
		want       int
	}{
		{[]int{1, 2, 3}, 2},
		{[]int{3, 3, 3, 4, 4, 4}, 0},
	}

	for _, tt := range tests {
		got := countPermutations(tt.complexity)
		if got != tt.want {
			t.Errorf("countPermutations(%v) = %d; want %d", tt.complexity, got, tt.want)
		}
	}
}
