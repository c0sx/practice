package leetcode757

import (
	"sort"
)

func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] > intervals[j][0]
		}

		return intervals[i][1] < intervals[j][1]
	})

	result := 0
	secondLast, last := -1, -1

	for _, interval := range intervals {
		start, end := interval[0], interval[1]

		if start <= secondLast {
			continue
		}

		if start > last {
			result += 2
			secondLast = end - 1
			last = end
		} else {
			result++
			secondLast = last
			last = end
		}
	}

	return result
}
