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

func balanceBST(root *TreeNode) *TreeNode {
	var list []int

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		list = append(list, root.Val)
		dfs(root.Right)
	}

	dfs(root)

	var build func(l, r int) *TreeNode
	build = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}

		m := (l + r) / 2
		node := &TreeNode{
			Val: list[m],
		}

		node.Left = build(l, m-1)
		node.Right = build(m+1, r)

		return node
	}

	tree := build(0, len(list)-1)

	return tree
}
