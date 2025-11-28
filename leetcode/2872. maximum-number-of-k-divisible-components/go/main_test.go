package leetcode2872

import "testing"

func TestMaxKDivisibleComponents(t *testing.T) {
	tests := []struct {
		n      int
		edges  [][]int
		values []int
		k      int
		expect int
	}{
		{
			n:      5,
			edges:  [][]int{{0, 2}, {1, 2}, {1, 3}, {2, 4}},
			values: []int{1,8,1,4,4},
			k:      6,
			expect: 2,
		},
		{
			n:      7,
			edges:  [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}},
			values: []int{3, 0, 6, 1,  5, 2, 1},
			k:      3,
			expect: 3,
		},
	}

	for _, test := range tests {
		result := maxKDivisibleComponents(test.n, test.edges, test.values, test.k)
		if result != test.expect {
			t.Errorf("For n=%d, edges=%v, values=%v, k=%d: expected %d, but got %d",
				test.n, test.edges, test.values, test.k, test.expect, result)
		}
	}
}
