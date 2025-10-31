package leetcode3289

func getSneakyNumbers(nums []int) []int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	result := []int{}
	for num, count := range counter {
		if count == 2 {
			result = append(result, num)
		}
	}

	return result
}
