package leetcode1161

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	counters := make(map[int]int)

	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		counters[level] += node.Val
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}

	dfs(root, 1)

	result := -1
	maxSum := 0
	for level, sum := range counters {
		if result == -1 || sum > maxSum {
			result = level
			maxSum = sum
		}
	}

	return result
}
