package leetcode961

func repeatedNTimes(nums []int) int {
	counters := make(map[int]int)

	for _, num := range nums {
		counters[num]++

		if counters[num] > 1 {
			return num
		}
	}

	return -1
}
