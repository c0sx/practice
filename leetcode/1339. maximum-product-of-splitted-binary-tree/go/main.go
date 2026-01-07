package leetcode1339

const MOD = 1e9 + 7

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxProduct(root *TreeNode) int {
	total, result := 0, 0

	var sum func(node *TreeNode) int
	sum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		return node.Val + sum(node.Left) + sum(node.Right)
	}

	total = sum(root)

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		s := node.Val + dfs(node.Left) + dfs(node.Right)
		prod := s * (total - s)
		if prod > result {
			result = prod
		}

		return s
	}

	dfs(root)

	return result % MOD
}
