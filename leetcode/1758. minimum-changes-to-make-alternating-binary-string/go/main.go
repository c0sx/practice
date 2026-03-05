package lc

func minOperations(s string) int {
	start0 := 0
	start1 := 0

	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			if s[i] == '0' {
				start1++
			} else {
				start0++
			}
		} else {
			if s[i] == '1' {
				start1++
			} else {
				start0++
			}
		}
	}

	return min(start0, start1)
}
