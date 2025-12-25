package leetcode3075

import "testing"

func TestMaximumHappinessSum(t *testing.T) {
	tests := []struct {
		happiness []int
		k         int
		expected  int64
	}{
		{happiness: []int{1, 2, 3}, k: 2, expected: 4},
		{happiness: []int{1, 1, 1, 1}, k: 2, expected: 1},
		{happiness: []int{2, 3, 4, 5}, k: 1, expected: 5},
		{happiness: []int{12, 1, 42}, k: 3, expected: 53},
	}

	for _, test := range tests {
		result := maximumHappinessSum(test.happiness, test.k)
		if result != test.expected {
			t.Errorf("For happiness %v and k %d, expected %d but got %d", test.happiness, test.k, test.expected, result)
		}
	}
}
