package leetcode3005

func maxFrequencyElements(nums []int) int {
	freqMap := make(map[int]int)

	for _, num := range nums {
		freqMap[num]++
	}

	maxFreq := 0
	for _, freq := range freqMap {
		if freq > maxFreq {
			maxFreq = freq
		}
	}

	count := 0
	for _, freq := range freqMap {
		if freq == maxFreq {
			count += freq
		}
	}

	return count
}
