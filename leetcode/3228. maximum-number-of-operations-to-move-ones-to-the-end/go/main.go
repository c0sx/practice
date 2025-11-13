package leetcode3228

func maxOperations(s string) int {
	n := len(s)
	counter := 0
	ones := 0

	i := 0
	for i < n {
		if s[i] == '0' {
			for i+1 < n && s[i+1] == '0' {
				i++
			}

			counter += ones
		} else {
			ones++
		}

		i++
	}

	return counter
}
