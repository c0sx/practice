package leetcode1339

import "testing"

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "test1",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
				},
			},
			want: 110,
		},
		{
			name: "test2",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
			},
			want: 90,
		},
	}

	for _, tt := range tests {
		if got := maxProduct(tt.root); got != tt.want {
			t.Errorf("maxProduct(%v) = %v, want %v", tt.root, got, tt.want)
		}
	}
}
