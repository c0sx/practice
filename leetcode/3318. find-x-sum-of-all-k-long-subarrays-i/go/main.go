package leetcode3318

import "sort"

type Pair struct {
	freq  int
	value int
}

func findXSum(nums []int, k int, x int) []int {
	n := len(nums)
	result := make([]int, 0, n-k+1)

	for i := 0; i <= n-k; i++ {
		counter := make(map[int]int)

		for j := i; j < i+k; j++ {
			counter[nums[j]]++
		}

		freq := make([]Pair, 0, len(counter))
		for val, count := range counter {
			freq = append(freq, Pair{count, val})
		}

		sort.Slice(freq, func(i, j int) bool {
			if freq[i].freq != freq[j].freq {
				return freq[i].freq > freq[j].freq
			}

			return freq[i].value > freq[j].value
		})

		xsum := 0
		for j := 0; j < x && j < len(freq); j++ {
			xsum += freq[j].freq * freq[j].value
		}

		result = append(result, xsum)
	}

	return result
}
