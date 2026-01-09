package leetcode865

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
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode, depth int) (subtree *TreeNode, maxDepth int)
	dfs = func(node *TreeNode, depth int) (subtree *TreeNode, maxDepth int) {
		if node == nil {
			return nil, depth
		}

		leftSubtree, leftMaxDepth := dfs(node.Left, depth+1)
		rightSubtree, rightMaxDepth := dfs(node.Right, depth+1)

		if leftMaxDepth > rightMaxDepth {
			return leftSubtree, leftMaxDepth
		} else if rightMaxDepth > leftMaxDepth {
			return rightSubtree, rightMaxDepth
		} else {
			return node, leftMaxDepth
		}
	}

	subtree, _ := dfs(root, 0)

	return subtree
}
