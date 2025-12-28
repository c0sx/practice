package leetcode1351

import "testing"

/**
Example 1:

Input: grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
Output: 8
Explanation: There are 8 negatives number in the matrix.
Example 2:

Input: grid = [[3,2],[1,0]]
Output: 0
*/

func TestCountNegatives(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}, 8},
		{[][]int{{3, 2}, {1, 0}}, 0},
	}

	for _, tt := range tests {
		got := countNegatives(tt.grid)
		if got != tt.want {
			t.Errorf("countNegatives(%v) = %v; want %v", tt.grid, got, tt.want)
		}
	}
}
