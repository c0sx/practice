package leetcode3047

func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
	n := len(bottomLeft)
	maxSide := int64(0)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			width := min(topRight[i][0], topRight[j][0]) - max(bottomLeft[i][0], bottomLeft[j][0])
			height := min(topRight[i][1], topRight[j][1]) - max(bottomLeft[i][1], bottomLeft[j][1])

			if width > 0 && height > 0 {
				side := int64(min(width, height))

				if side > maxSide {
					maxSide = side
				}
			}
		}
	}

	return maxSide * maxSide
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
