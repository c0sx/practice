package leetcode3432

func countPartitions(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	if sum%2 == 0 {
		return len(nums) - 1
	}

	return 0
}
