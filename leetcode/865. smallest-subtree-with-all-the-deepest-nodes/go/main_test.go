package leetcode865

import "testing"

func TestSubtreeWithAllDeepest(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want *TreeNode
	}{
		{
			name: "test1",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
				},
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
			want: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		{
			name: "test2",
			root: &TreeNode{
				Val: 1,
			},
			want: &TreeNode{
				Val: 1,
			},
		},
		{
			name: "test3",
			root: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: &TreeNode{
				Val: 2,
			},
		},
	}

	for _, tt := range tests {
		if got := subtreeWithAllDeepest(tt.root); !equal(got, tt.want) {
			t.Errorf("subtreeWithAllDeepest() = %v, want %v", got, tt.want)
		}
	}
}

func equal(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Val != b.Val {
		return false
	}

	return equal(a.Left, b.Left) && equal(a.Right, b.Right)
}
