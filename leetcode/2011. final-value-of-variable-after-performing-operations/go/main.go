package leetcode2011

func finalValueAfterOperations(operations []string) int {
	x := 0

	for _, op := range operations {
		if op[0] == '+' || op[len(op)-1] == '+' {
			x++
		} else {
			x--
		}
	}

	return x
}
