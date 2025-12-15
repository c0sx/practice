package leetcode2110

func getDescentPeriods(prices []int) int64 {
	n := len(prices)
	var result int64 = 1
	prev := 1

	for i := 1; i < n; i++ {
		if prices[i] == prices[i-1]-1 {
			prev++
		} else {
			prev = 1
		}

		result += int64(prev)
	}

	return result
}
