package lc

func longestBalanced(s string) int {
	maxLen := 0
	n := len(s)

	for i := 0; i < n; i++ {
		chars := make(map[byte]int)

		for j := i; j < n; j++ {
			c := s[j]
			chars[c]++

			found := true
			var val int = 0
			for _, v := range chars {
				if val == 0 && v != 0 {
					val = v
					continue
				}

				if val != v {
					found = false
					break
				}
			}

			if found {
				len := j - i + 1
				maxLen = max(maxLen, len)
			}
		}
	}

	return maxLen
}
