package lc

import (
	"fmt"
	"testing"
)

func TestBalanceBST(t *testing.T) {
	tests := []struct {
		root *TreeNode
	}{
		{
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 4,
						},
					},
				},
			},
		},
		// {
		// 	root: &TreeNode{
		// 		Val: 2,
		// 		Left: &TreeNode{
		// 			Val: 1,
		// 		},
		// 		Right: &TreeNode{
		// 			Val: 3,
		// 		},
		// 	},
		// },
		// {
		// 	root: &TreeNode{
		// 		Val: 1,
		// 		Right: &TreeNode{
		// 			Val: 2,
		// 			Right: &TreeNode{
		// 				Val: 3,
		// 				Right: &TreeNode{
		// 					Val: 4,
		// 					Right: &TreeNode{
		// 						Val: 1,
		// 						Right: &TreeNode{
		// 							Val: 2,
		// 							Right: &TreeNode{
		// 								Val: 3,
		// 								Right: &TreeNode{
		// 									Val: 4,
		// 								},
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// },
	}

	for _, tt := range tests {
		got := balanceBST(tt.root)

		var print func(root *TreeNode, level int)
		print = func(root *TreeNode, level int) {
			if root == nil {
				return
			}

			tab := ""
			for i := 0; i < level; i++ {
				tab += " "
			}

			fmt.Printf("%s%d\n", tab, root.Val)

			print(root.Left, level+1)
			print(root.Right, level+1)
		}

		print(got, 0)

		if !isBalanced(got) {
			t.Errorf("balanceBST is not balanced")
		}
	}
}

func isBalanced(root *TreeNode) bool {
	var solve func(root *TreeNode) int
	solve = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := solve(root.Left)
		if left == -1 {
			return -1
		}

		right := solve(root.Right)
		if right == -1 {
			return -1
		}

		diff := abs(left - right)
		if diff > 1 {
			return -1
		}

		if left > right {
			return 1 + left
		}

		return 1 + right
	}

	return solve(root) != -1
}

func abs(val int) int {
	if val < 0 {
		return -val
	}

	return val
}
