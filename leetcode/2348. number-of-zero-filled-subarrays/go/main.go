package leetcode2348

func zeroFilledSubarray(nums []int) int64 {
	total := int64(0)

	counter := int64(0)
	for _, num := range nums {
		if num == 0 {
			counter++
		} else {
			counter = 0
		}

		total += counter
	}

	return total
}
