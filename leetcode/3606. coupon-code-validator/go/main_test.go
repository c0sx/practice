package leetcode3606

import "testing"

func TestValidateCoupons(t *testing.T) {
	tests := []struct {
		code         []string
		businessLine []string
		isActive     []bool
		expected     []string
	}{
		{
			code:         []string{"SAVE20", "", "PHARMA5", "SAVE@20"},
			businessLine: []string{"restaurant", "grocery", "pharmacy", "restaurant"},
			isActive:     []bool{true, true, true, true},
			expected:     []string{"PHARMA5", "SAVE20"},
		},
		{
			code:         []string{"GROCERY15", "ELECTRONICS_50", "DISCOUNT10"},
			businessLine: []string{"grocery", "electronics", "invalid"},
			isActive:     []bool{false, true, true},
			expected:     []string{"ELECTRONICS_50"},
		},
		{
			code:         []string{"rm"},
			businessLine: []string{"pharmacy"},
			isActive:     []bool{true},
			expected:     []string{"rm"},
		},
		{
			code:         []string{"TsCwKhY", "qCeVkfb", "ZGjX07H"},
			businessLine: []string{"restaurant", "electronics", "pharmacy"},
			isActive:     []bool{true, true, true},
			expected:     []string{"qCeVkfb", "ZGjX07H", "TsCwKhY"},
		},
	}

	for _, test := range tests {
		result := validateCoupons(test.code, test.businessLine, test.isActive)
		if len(result) != len(test.expected) {
			t.Errorf("For input code: %v, businessLine: %v, isActive: %v, expected length %d but got %d", test.code, test.businessLine, test.isActive, len(test.expected), len(result))
			continue
		}

		if len(result) != len(test.expected) {
			t.Errorf("For input code: %v, businessLine: %v, isActive: %v, expected length %d but got %d", test.code, test.businessLine, test.isActive, len(test.expected), len(result))
			continue
		}

		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("For input code: %v, businessLine: %v, isActive: %v, expected %v but got %v", test.code, test.businessLine, test.isActive, test.expected, result)
				break
			}
		}
	}
}
