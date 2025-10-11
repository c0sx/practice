package leetcode3186

import "sort"

func maximumTotalDamage(power []int) int64 {
	count := map[int]int{}
	for _, p := range power {
		count[p]++
	}

	keys := make([]int, 0, len(count))
	for k := range count {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	vec := [][2]int{}
	for _, k := range keys {
		vec = append(vec, [2]int{k, count[k]})
	}

	n := len(vec)
	f := make([]int64, n)

	var max int64
	var result int64

	j := 0
	for i := 0; i < n; i++ {
		for j < i && vec[j][0] < vec[i][0]-2 {
			if f[j] > max {
				max = f[j]
			}

			j++
		}

		f[i] = max + int64(vec[i][0])*int64(vec[i][1])
		if f[i] > result {
			result = f[i]
		}
	}

	return result
}
