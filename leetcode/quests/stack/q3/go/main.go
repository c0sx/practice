package q3

import "strings"

func exclusiveTime(n int, logs []string) []int {
	result := make([]int, n)
	s := []int{}
	prev := 0

	for i, log := range logs {
		parts := strings.Split(log, ":")
		id := toInt(parts[0])
		action := parts[1]
		timestamp := toInt(parts[2])

		if i == 0 {
			s = append(s, id)
			prev = timestamp
			continue
		}

		switch action {
		case "start":
			if len(s) > 0 {
				top := s[len(s)-1]
				result[top] += timestamp - prev
			}

			s = append(s, id)
			prev = timestamp
		case "end":
			top := s[len(s)-1]
			s = s[:len(s)-1]
			result[top] += timestamp - prev + 1
			prev = timestamp + 1
		}
	}

	return result
}

func toInt(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res = res*10 + int(s[i]-'0')
	}

	return res
}
