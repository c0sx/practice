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

func sumRootToLeaf(root *TreeNode) int {
	var result int

	var solve func(root *TreeNode, prev uint64)
	solve = func(root *TreeNode, prev uint64) {
		if root == nil {
			return
		}

		b := uint64(root.Val)
		prev += b

		if root.Left == nil && root.Right == nil {
			result += int(prev)
		} else {
			prev = prev << 1
			solve(root.Left, prev)
			solve(root.Right, prev)
		}
	}

	solve(root, 0)

	return result
}
