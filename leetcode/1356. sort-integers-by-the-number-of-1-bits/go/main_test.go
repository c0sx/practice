package lc

import "testing"

func TestSortByBits(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			want: []int{0, 1, 2, 4, 8, 3, 5, 6, 7},
		},
		{
			arr:  []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1},
			want: []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024},
		},
	}

	for _, tt := range tests {
		got := sortByBits(tt.arr)
		if !equal(got, tt.want) {
			t.Errorf("sortByBits(%v) = %v, want %v", tt.arr, got, tt.want)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
