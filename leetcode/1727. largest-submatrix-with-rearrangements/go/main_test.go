package lc

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   int
	}{
		{
			matrix: [][]int{{0, 0, 1}, {1, 1, 1}, {1, 0, 1}},
			want:   4,
		},
		{
			matrix: [][]int{{1, 0, 1, 0, 1}},
			want:   3,
		},
		{
			matrix: [][]int{{1, 1, 0}, {1, 0, 1}},
			want:   2,
		},
		{
			matrix: [][]int{
				{1, 1, 0, 0, 1, 0, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1},
				{1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 1},
				{1, 1, 1, 0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 0},
			},
			want: 40,
		},
	}

	for _, tt := range tests {
		got := largestSubmatrix(tt.matrix)
		if got != tt.want {
			t.Errorf("%v = %d, want %d", tt.matrix, got, tt.want)
		}
	}
}
