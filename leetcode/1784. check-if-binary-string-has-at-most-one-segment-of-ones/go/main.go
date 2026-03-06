package lc

func checkOnesSegment(s string) bool {
	n := len(s)
	flag := false

	for i := 1; i < n; i++ {
		if s[i] == '0' {
			flag = true
		} else {
			if flag {
				return false
			}
		}
	}

	return true
}
