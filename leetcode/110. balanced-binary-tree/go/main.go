package lc

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
