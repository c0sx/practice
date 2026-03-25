package lc

func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := make([]int, 0, n)
	maxArea := 0

	for i := 0; i <= len(heights); i++ {
		height := 0
		if i < len(heights) {
			height = heights[i]
		}

		for len(stack) > 0 && heights[stack[len(stack)-1]] > height {
			top := len(stack) - 1
			mid := stack[top]
			stack = stack[:top]

			left := -1
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}

			width := i - left - 1
			area := heights[mid] * width
			if area > maxArea {
				maxArea = area
			}
		}

		stack = append(stack, i)
	}

	return maxArea
}
