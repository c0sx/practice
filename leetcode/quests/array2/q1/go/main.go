package q1

func findErrorNums(nums []int) []int {
	counter := make(map[int]bool)
	n := len(nums)

	actualSum := (n * (n + 1)) / 2
	result := []int{}

	for _, num := range nums {
		if _, exists := counter[num]; exists {
			result = append(result, num)
		}

		counter[num] = true
	}

	sum := 0
	for num, _ := range counter {
		sum += num
	}

	missing := actualSum - sum
	result = append(result, missing)

	return result
}
