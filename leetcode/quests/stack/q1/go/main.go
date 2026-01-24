package q1

func buildArray(target []int, n int) []string {
	ops := []string{}
	s := []int{}
	j := 0

	for i := 1; i <= n; i++ {
		s = append(s, i)
		ops = append(ops, "Push")

		if target[j] != i {
			s = s[:len(s)-1]
			ops = append(ops, "Pop")
		} else {
			j++
		}

		if j >= len(target) {
			break
		}
	}

	return ops
}
