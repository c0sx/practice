package leetcode1625

func findLexSmallestString(s string, a int, b int) string {
	queue := []string{s}
	visited := map[string]bool{s: true}
	min := s

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current < min {
			min = current
		}

		temp1 := []byte(current)
		for i := 1; i < len(temp1); i += 2 {
			c := int(temp1[i]-'0') + a
			temp1[i] = byte((c % 10) + '0')
		}

		s1 := string(temp1)
		s2 := current[b:] + current[:b]
		candidates := []string{s1, s2}

		for _, candidate := range candidates {
			if !visited[candidate] {
				visited[candidate] = true
				queue = append(queue, candidate)
			}
		}
	}

	return min
}
