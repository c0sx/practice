package leetcode417

import (
	"testing"
)

func TestPacificAtlantic(t *testing.T) {
	tests := []struct {
		input [][]int
		want  [][]int
	}{
		{
			input: [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}},
			want:  [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}},
		},
		{
			input: [][]int{{1}},
			want:  [][]int{{0, 0}},
		},
	}

	for _, tt := range tests {
		result := pacificAtlantic(tt.input)
		if !equal(result, tt.want) {
			t.Errorf("pacificAtlantic(%v) = %v, want %v", tt.input, result, tt.want)
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
