package leetcode1590

func minSubarray(nums []int, p int) int {
	n := len(nums)
	totalSum := 0

	for _, num := range nums {
		totalSum = (totalSum + num) % p
	}

	target := totalSum % p
	if target == 0 {
		return 0
	}

	modMap := map[int]int{0: -1}
	currentSum := 0
	minLength := n

	for i := 0; i < n; i++ {
		currentSum = (currentSum + nums[i]) % p
		needed := (currentSum - target + p) % p

		if idx, exists := modMap[needed]; exists {
			minLength = min(minLength, i-idx)
		}

		modMap[currentSum] = i
	}

	if minLength < n {
		return minLength
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
