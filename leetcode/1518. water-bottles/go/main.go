package leetcode1518

func numWaterBottles(numBottles int, numExchange int) int {
	drunk := 0

	for numBottles > 0 {
		numBottles--
		drunk++

		if drunk%numExchange == 0 {
			numBottles++
		}
	}

	return drunk
}
