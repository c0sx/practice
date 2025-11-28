package leetcode2872

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	adj := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	var dfs func(node, parent int, count *int) int
	dfs = func(node, parent int, count *int) int {
		sum := 0

		for _, neighbor := range adj[node] {
			if neighbor != parent {
				sum += dfs(neighbor, node, count)
				sum %= k
			}
		}

		sum += values[node]
		sum %= k

		if sum == 0 {
			*count++
		}

		return sum
	}

	count := 0
	dfs(0, -1, &count)

	return count
}
