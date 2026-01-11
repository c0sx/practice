package leetcode85

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	m := len(matrix[0])

	heights := make([]int, m)
	maxArea := 0

	for _, row := range matrix {
		for j := 0; j < m; j++ {
			if row[j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}

		maxArea = max(maxArea, largestRectangleArea(heights))
	}

	return maxArea
}

func largestRectangleArea(heights []int) int {
	stack := []int{}
	heights = append(heights, 0)
	area := 0

	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i

			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			area = max(area, height*width)
		}

		stack = append(stack, i)
	}

	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
