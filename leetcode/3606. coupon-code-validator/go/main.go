package leetcode3606

import (
	"sort"
)

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	businessLineMap := map[string]int{
		"restaurant":  4,
		"grocery":     2,
		"pharmacy":    3,
		"electronics": 1,
	}

	validCoupons := []int{}

	for i := 0; i < len(code); i++ {
		if isActive[i] && businessLineMap[businessLine[i]] > 0 && isValidCode(code[i]) {
			validCoupons = append(validCoupons, i)
		}
	}

	sort.Slice(validCoupons, func(i, j int) bool {
		businessLineA := businessLine[validCoupons[i]]
		businessLineB := businessLine[validCoupons[j]]

		priorityA := businessLineMap[businessLineA]
		priorityB := businessLineMap[businessLineB]

		if priorityA != priorityB {
			return priorityA < priorityB
		}

		codeA := code[validCoupons[i]]
		codeB := code[validCoupons[j]]

		return codeA < codeB
	})

	result := []string{}
	for _, idx := range validCoupons {
		result = append(result, code[idx])
	}

	return result
}

func isValidCode(coupon string) bool {
	if len(coupon) == 0 {
		return false
	}

	for i := 0; i < len(coupon); i++ {
		ch := coupon[i]
		if !(ch >= 'a' && ch <= 'z') &&
			!(ch >= 'A' && ch <= 'Z') &&
			!(ch >= '0' && ch <= '9') &&
			ch != '_' {
			return false
		}
	}

	return true
}
