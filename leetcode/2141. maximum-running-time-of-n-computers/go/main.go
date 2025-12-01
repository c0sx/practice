package leetcode2141

import "sort"

func maxRunTime(n int, batteries []int) int64 {
	var totalEnergy int64 = 0
	for _, b := range batteries {
		totalEnergy += int64(b)
	}

	sort.Ints(batteries)

	for i := len(batteries) - 1; i >= 0; i-- {
		if int64(batteries[i]) > totalEnergy/int64(n) {
			totalEnergy -= int64(batteries[i])
			n--
		} else {
			break
		}
	}

	return totalEnergy / int64(n)
}
