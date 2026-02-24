package lc

import "testing"

func TestSumRootToLeaf(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
			want: 22,
		},
		{
			root: &TreeNode{
				Val: 0,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		got := sumRootToLeaf(tt.root)
		if got != tt.want {
			t.Errorf("sumRootToLeaf(%v) = %d, want %d", tt.root, got, tt.want)
		}
	}
}
