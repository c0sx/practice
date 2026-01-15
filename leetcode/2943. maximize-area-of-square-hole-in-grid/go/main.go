package leetcode2943

import "sort"

func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	sort.Ints(hBars)
	sort.Ints(vBars)

	hMax, vMax := 1, 1
	hCur, vCur := 1, 1

	for i := 1; i < len(hBars); i++ {
		if hBars[i] == hBars[i-1]+1 {
			hCur++
		} else {
			hCur = 1
		}

		hMax = max(hMax, hCur)
	}

	for i := 1; i < len(vBars); i++ {
		if vBars[i] == vBars[i-1]+1 {
			vCur++
		} else {
			vCur = 1
		}

		vMax = max(vMax, vCur)
	}

	side := min(hMax, vMax) + 1

	return side * side
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
