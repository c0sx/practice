package q2

func evalRPN(tokens []string) int {
	s := []int{}

	for _, token := range tokens {
		switch token {
		case "+":
			n2, n1 := s[len(s)-1], s[len(s)-2]
			s = s[:len(s)-2]
			s = append(s, n1+n2)
		case "-":
			n2, n1 := s[len(s)-1], s[len(s)-2]
			s = s[:len(s)-2]
			s = append(s, n1-n2)
		case "*":
			n2, n1 := s[len(s)-1], s[len(s)-2]
			s = s[:len(s)-2]
			s = append(s, n1*n2)
		case "/":
			n2, n1 := s[len(s)-1], s[len(s)-2]
			s = s[:len(s)-2]
			s = append(s, n1/n2)
		default:
			var num int
			for i := 0; i < len(token); i++ {
				if i == 0 && token[i] == '-' {
					continue
				}

				num = num*10 + int(token[i]-'0')
			}

			if token[0] == '-' {
				num = -num
			}

			s = append(s, num)
		}
	}

	return s[0]
}
