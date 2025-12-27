package leetcode2402

import "sort"

func mostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	count := make([]int, n)
	timer := make([]int64, n)

	itr := 0

	for itr < len(meetings) {
		start := meetings[itr][0]
		end := meetings[itr][1]
		dur := int64(end - start)

		room := -1
		earliest := int64(1 << 62)
		earliestRoom := -1

		for i := 0; i < n; i++ {
			if timer[i] < earliest {
				earliest = timer[i]
				earliestRoom = i
			}
			if timer[i] <= int64(start) {
				room = i
				break
			}
		}

		if room != -1 {
			timer[room] = int64(end)
			count[room]++
		} else {
			timer[earliestRoom] += dur
			count[earliestRoom]++
		}

		itr++
	}

	max := 0
	idx := 0
	for i := 0; i < n; i++ {
		if count[i] > max {
			max = count[i]
			idx = i
		}
	}

	return idx
}
