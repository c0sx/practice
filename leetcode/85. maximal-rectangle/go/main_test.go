package leetcode85

import "testing"

func TestMaximalRectangle(t *testing.T) {
	tests := []struct {
		matrix [][]byte
		expect int
	}{
		{
			matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			expect: 6,
		},
		{
			matrix: [][]byte{},
			expect: 0,
		},
		{
			matrix: [][]byte{
				{'0'},
			},
			expect: 0,
		},
		{
			matrix: [][]byte{
				{'1'},
			},
			expect: 1,
		},
	}

	for _, test := range tests {
		result := maximalRectangle(test.matrix)
		if result != test.expect {
			t.Errorf("For matrix %v, expected %d but got %d", test.matrix, test.expect, result)
		}
	}
}
