package leetcode1161

import "testing"

/**
Input: root = [1,7,0,7,-8,null,null]
Output: 2

Input: root = [989,null,10250,98693,-89388,null,null,null,-32127]
Output: 2
*/

func TestMaxLevelSum(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "test case 1",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: -8,
					},
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
			want: 2,
		},
		{
			name: "test case 2",
			root: &TreeNode{
				Val: 989,
				Right: &TreeNode{
					Val: 10250,
					Left: &TreeNode{
						Val: 98693,
					},
					Right: &TreeNode{
						Val: -89388,
						Right: &TreeNode{
							Val: -32127,
						},
					},
				},
			},
			want: 2,
		},
	}
	
	for _, tt := range tests {
		if got := maxLevelSum(tt.root); got != tt.want {
			t.Errorf("maxLevelSum(%v) = %v, want %v", tt.root, got, tt.want)
		}
	}
}
