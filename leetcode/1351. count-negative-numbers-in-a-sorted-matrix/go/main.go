package leetcode1351

func countNegatives(grid [][]int) int {
	count := 0

	for _, row := range grid {
		n := len(row)

		l := 0
		r := n - 1

		for l <= r {
			mid := l + (r-l)/2
			if row[mid] < 0 {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

		count += n - l
	}

	return count
}
