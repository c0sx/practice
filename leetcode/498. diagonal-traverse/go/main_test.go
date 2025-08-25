package leetcode498

import "testing"

func TestFindDiagonalOrder(t *testing.T) {
	tests := []struct {
		mat  [][]int
		want []int
	}{
		
		{
			mat:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want: []int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
		// {
		// 	mat:  [][]int{{1, 2}, {3, 4}},
		// 	want: []int{1, 2, 3, 4},
		// },
		// {
		// 	mat:  [][]int{{1}},
		// 	want: []int{1},
		// },
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findDiagonalOrder(tt.mat); !equal(got, tt.want) {
				t.Errorf("findDiagonalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
