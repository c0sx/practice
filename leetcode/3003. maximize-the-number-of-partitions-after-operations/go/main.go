package leetcode3003

func maxPartitionsAfterOperations(s string, k int) int {
	n := len(s)
	memo := make(map[int64]int)

	var dfs func(int, int, int) int
	dfs = func(index, mask, canChange int) int {
		if index >= n {
			if mask == 0 {
				return 0
			}

			return 1
		}

		key := (int64(index) << 27) | (int64(mask) << 1) | int64(canChange)
		if v, ok := memo[key]; ok {
			return v
		}

		bit := 1 << (s[index] - 'a')
		nextMask := mask | bit

		var ans int
		if bitCount(nextMask) > k {
			ans = dfs(index+1, bit, canChange) + 1
		} else {
			ans = dfs(index+1, nextMask, canChange)
		}

		if canChange == 1 {
			for c := 0; c < 26; c++ {
				if c == int(s[index]-'a') {
					continue
				}
				bitC := 1 << c
				newMask := mask | bitC
				if bitCount(newMask) > k {
					ans = max(ans, dfs(index+1, bitC, 0)+1)
				} else {
					ans = max(ans, dfs(index+1, newMask, 0))
				}
			}
		}

		memo[key] = ans
		return ans
	}

	return dfs(0, 0, 1)
}

func bitCount(x int) int {
	x = x - ((x >> 1) & 0x55555555)
	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)

	return x & 0x3f
}

func max(a, b int) int {
	if a > b {
		return a
	}
	
	return b
}
