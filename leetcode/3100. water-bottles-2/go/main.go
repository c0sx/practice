package leetcode3100

func maxBottlesDrunk(numBottles int, numExchange int) int {
	drunk := 0
	emptyBottles := 0

	for numBottles > 0 {
		emptyBottles += numBottles
		drunk += numBottles
		numBottles = 0

		for emptyBottles >= numExchange {
			emptyBottles -= numExchange
			numBottles++
			numExchange++
		}
	}

	return drunk
}
