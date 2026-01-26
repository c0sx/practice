package leetcode1200

import "testing"

/**
Example 1:

Input: arr = [4,2,1,3]
Output: [[1,2],[2,3],[3,4]]

Example 2:

Input: arr = [1,3,6,10,15]
Output: [[1,3]]

Example 3:

Input: arr = [3,8,-10,23,19,-4,-14,27]
Output: [[-14,-10],[19,23],[23,27]]
*/

func Test_minimumAbsDifference(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want [][]int
	}{
		{
			name: "Example 1",
			arr:  []int{4, 2, 1, 3},
			want: [][]int{{1, 2}, {2, 3}, {3, 4}},
		},
		{
			name: "Example 2",
			arr:  []int{1, 3, 6, 10, 15},
			want: [][]int{{1, 3}},
		},
		{
			name: "Example 3",
			arr:  []int{3, 8, -10, 23, 19, -4, -14, 27},
			want: [][]int{{-14, -10}, {19, 23}, {23, 27}},
		},
	}

	for _, tt := range tests {
		if got := minimumAbsDifference(tt.arr); !equal(got, tt.want) {
			t.Errorf("minimumAbsDifference(%v) = %v, want %v", tt.arr, got, tt.want)
		}
	}
}

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}

		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}
