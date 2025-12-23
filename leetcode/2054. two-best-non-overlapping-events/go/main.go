package leetcode2054

import "sort"

func maxTwoEvents(events [][]int) int {
	n := len(events)
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, 3)

		for j := 0; j < 3; j++ {
			dp[i][j] = -1
		}
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	var findEvents func(events [][]int, dp [][]int, index int, count int) int
	findEvents = func(events [][]int, dp [][]int, index int, count int) int {
		if index >= len(events) || count >= 2 {
			return 0
		}

		if dp[index][count] == -1 {
			end := events[index][1]
			lo := index + 1
			hi := len(events) - 1

			for lo <= hi {
				mid := lo + (hi-lo)/2

				if events[mid][0] > end {
					hi = mid - 1
				} else {
					lo = mid + 1
				}
			}

			include := events[index][2] + findEvents(events, dp, lo, count+1)
			exclude := findEvents(events, dp, index+1, count)
			dp[index][count] = max(include, exclude)
		}

		return dp[index][count]
	}

	result := findEvents(events, dp, 0, 0)

	return result
}
