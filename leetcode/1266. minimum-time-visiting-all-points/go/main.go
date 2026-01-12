package leetcode1266

func minTimeToVisitAllPoints(points [][]int) int {
	result := 0
	n := len(points)

	for i := 0; i < n-1; i++ {
		x := points[i][0]
		y := points[i][1]

		nextX := points[i+1][0]
		nextY := points[i+1][1]

		result += max(abs(nextX-x), abs(nextY-y))
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
