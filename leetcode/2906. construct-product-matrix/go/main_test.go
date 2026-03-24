package lc

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		grid [][]int
		want [][]int
	}{
		{
			grid: [][]int{{1, 2}, {3, 4}},
			want: [][]int{{24, 12}, {8, 6}},
		},
		{
			grid: [][]int{{12345}, {2}, {1}},
			want: [][]int{{2}, {0}, {0}},
		},
		{
			grid: [][]int{{10, 20}, {18, 16}, {17, 14}, {16, 9}, {14, 6}, {16, 5}, {14, 8}, {20, 13}, {16, 10}, {14, 17}},
			want: [][]int{{345, 6345}, {7050, 4845}, {4560, 2010}, {4845, 1755}, {2010, 8805}, {4845, 690}, {2010, 9690}, {6345, 1215}, {4845, 345}, {2010, 4560}},
		},
	}

	for _, tt := range tests {
		got := constructProductMatrix(tt.grid)
		if !equal(got, tt.want) {
			t.Errorf("fn(%v) = %v, want %v", tt.grid, got, tt.want)
		}
	}
}

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}

		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}
