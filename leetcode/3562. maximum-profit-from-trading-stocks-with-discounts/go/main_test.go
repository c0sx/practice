package leetcode3562

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		n         int
		present   []int
		future    []int
		hierarchy [][]int
		budget    int
		expected  int
	}{
		{2, []int{1, 2}, []int{4, 3}, [][]int{{1, 2}}, 3, 5},
		{2, []int{3, 4}, []int{5, 8}, [][]int{{1, 2}}, 4, 4},
		{3, []int{4, 6, 8}, []int{7, 9, 11}, [][]int{{1, 2}, {1, 3}}, 10, 10},
		{3, []int{5, 2, 3}, []int{8, 5, 6}, [][]int{{1, 2}, {2, 3}}, 7, 12},
	}

	for _, test := range tests {
		result := maxProfit(test.n, test.present, test.future, test.hierarchy, test.budget)
		if result != test.expected {
			t.Errorf("maxProfit(%d, %v, %v, %v, %d) = %d; expected %d",
				test.n, test.present, test.future, test.hierarchy, test.budget, result, test.expected)
		}
	}
}
