package leetcode2048

func nextBeautifulNumber(n int) int {
	for candidate := n + 1; ; candidate++ {
		counts := make([]int, 10)

		tmp := candidate
		for tmp > 0 {
			lastDigit := tmp % 10
			counts[lastDigit]++
			tmp /= 10
		}

		isBeautiful := true
		for digit, count := range counts {
			if count != 0 && count != digit {
				isBeautiful = false
				break
			}
		}

		if isBeautiful {
			return candidate
		}
	}
}
