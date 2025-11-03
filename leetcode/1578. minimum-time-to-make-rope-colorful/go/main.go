package leetcode1578

func minCost(colors string, neededTime []int) int {
	total := 0
	n := len(colors)

	start := 0

	for start < n {
		end := start
		sum := 0
		maxTime := 0

		for end < n && colors[start] == colors[end] {
			sum += neededTime[end]
			maxTime = max(maxTime, neededTime[end])
			end++
		}

		if end-start > 1 {
			total += sum - maxTime
		}

		start = end
	}

	return total
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
