package q3

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	counters := make(map[int]bool)

	for _, num := range nums {
		counters[num] = true
	}

	result := []int{}
	for i := 1; i <= n; i++ {
		if _, exists := counters[i]; !exists {
			result = append(result, i)
		}
	}

	return result
}
