package leetcode166

import (
	"strconv"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	var result []string
	if (numerator < 0) != (denominator < 0) {
		result = append(result, "-")
	}

	dividend := abs(numerator)
	divisor := abs(denominator)

	result = append(result, strconv.Itoa(dividend/divisor))

	dividend = dividend % divisor
	if dividend == 0 {
		return strings.Join(result, "")
	}

	result = append(result, ".")
	remainders := make(map[int]int)

	for dividend != 0 {
		if idx, found := remainders[dividend]; found {
			result = append(result[:idx], "("+strings.Join(result[idx:], "")+")")
			break
		}

		remainders[dividend] = len(result)
		dividend *= 10
		result = append(result, strconv.Itoa(dividend/divisor))
		dividend = dividend % divisor

	}

	return strings.Join(result, "")
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
