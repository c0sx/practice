package leetcode3495

func minOperations(queries [][]int) int64 {
	var result int64

	for _, query := range queries {
		c1 := countOperations(query[1])
		c2 := countOperations(query[0] - 1)
		result += (c1 - c2 + 1) / 2
	}

	return result
}

func countOperations(n int) int64 {
	var counter int64
	base := 1
	i := 1

	for base <= n {
		end := base*2 - 1
		if end > n {
			end = n
		}

		counter += int64((i+1)/2) * int64(end-base+1)
		i++
		base *= 2
	}

	return counter
}
