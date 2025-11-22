package leetcode3190

func minimumOperations(nums []int) int {
	counter := 0

	for _, num := range nums {
		if num%3 != 0 {
			counter++
		}
	}

	return counter
}
